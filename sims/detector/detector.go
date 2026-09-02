// Copyright (c) 2026, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// detector: This simulation shows how an individual neuron can
// act like a detector, picking out specific patterns from its inputs
// and responding with varying degrees of selectivity to the match
// between its synaptic weights and the input activity pattern.
package detector

//go:generate core generate -add-types -add-funcs

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"reflect"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/base/metadata"
	"cogentcore.org/core/core"
	"cogentcore.org/core/enums"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/math32"
	"cogentcore.org/core/tree"
	"cogentcore.org/lab/base/mpi"
	"cogentcore.org/lab/base/randx"
	"cogentcore.org/lab/plot"
	"cogentcore.org/lab/stats/stats"
	"cogentcore.org/lab/table"
	"cogentcore.org/lab/tensor"
	"cogentcore.org/lab/tensorfs"
	"github.com/emer/axon/v2/axon"
	"github.com/emer/emergent/v2/egui"
	"github.com/emer/emergent/v2/env"
	"github.com/emer/emergent/v2/looper"
	"github.com/emer/emergent/v2/paths"
)

//go:embed digits.tsv
var embedfs embed.FS

// Modes are the looping modes (Stacks) for running and statistics.
type Modes int32 //enums:enum
const (
	Test  Modes = iota
	Train       // not used, but needed for some things
)

// Levels are the looping levels for running and statistics.
type Levels int32 //enums:enum
const (
	Cycle Levels = iota
	Trial
	Epoch
)

// DigitWts is the digit whose input pattern is copied into the receiving
// neuron's synaptic weights, thereby determining what it detects.
const DigitWts = 8

// see params.go for params, config.go for config

// Sim encapsulates the entire simulation model, and we define all the
// functionality as methods on this struct.  This structure keeps all relevant
// state information organized and available without having to pass everything around
// as arguments to methods, and provides the core GUI interface (note the view tags
// for the fields which provide hints to how things should be displayed).
type Sim struct {

	// GbarL is the leak conductance for the receiving neuron, in nS (nanosiemens),
	// which pulls against the excitatory input conductance to determine how hard
	// it is to activate the receiving neuron. This is much larger than the
	// default of 20 for a typical neuron, because there is no inhibition in this
	// network, so leak has to do all the work of counteracting excitation.
	// You must press Init after changing this for it to take effect!
	GbarL float32 `default:"240" min:"0" max:"400" step:"10"`

	// Config has simulation configuration parameters, set by .toml config file and / or args.
	Config *Config `new-window:"+"`

	// Net is the network: click to view / edit parameters for layers, paths, etc.
	Net *axon.Network `new-window:"+" display:"no-inline"`

	// Params manages network parameter setting.
	Params axon.Params `display:"inline"`

	// Loops are the control loops for running the sim, in different Modes
	// across stacks of Levels.
	Loops *looper.Stacks `new-window:"+" display:"no-inline"`

	// Envs provides mode-string based storage of environments.
	Envs env.Envs `new-window:"+" display:"no-inline"`

	// NetUpdate has netview update parameters.
	NetUpdate axon.NetViewUpdate `display:"inline"`

	// Root is the root tensorfs directory, where all stats and other misc sim data goes.
	Root *tensorfs.Node `display:"-"`

	// Stats has the stats directory within Root.
	Stats *tensorfs.Node `display:"-"`

	// Current has the current stats values within Stats.
	Current *tensorfs.Node `display:"-"`

	// StatFuncs are statistics functions called at given mode and level,
	// to perform all stats computations. phase = Start does init at start of given level,
	// and all intialization / configuration (called during Init too).
	StatFuncs []func(mode enums.Enum, level enums.Enum, start bool) `display:"-"`

	// GUI manages all the GUI elements
	GUI egui.GUI `display:"-"`

	// RandSeeds is a list of random seeds to use for each run.
	RandSeeds randx.Seeds `display:"-"`
}

func Embed(b tree.Node)               { egui.Embed[Sim, Config](b) }
func (ss *Sim) SetConfig(cfg *Config) { ss.Config = cfg }
func (ss *Sim) Body() *core.Body      { return ss.GUI.Body }

func (ss *Sim) Defaults() {
	ss.Config.Defaults()
	ss.GbarL = 240
}

func (ss *Sim) ConfigSim() {
	ss.Defaults()
	ss.Root, _ = tensorfs.NewDir("Root")
	tensorfs.CurRoot = ss.Root
	ss.Net = axon.NewNetwork(ss.Config.Name)
	ss.Params.Config(LayerParams, PathParams, ss.Config.Params.Sheet, ss.Config.Params.Tag, reflect.ValueOf(ss))
	ss.RandSeeds.Init(100) // max 100 runs
	ss.InitRandSeed(0)
	ss.OpenInputs()
	ss.ConfigEnv()
	ss.ConfigNet(ss.Net)
	ss.ConfigLoops()
	ss.ConfigStats()
	if ss.Config.Params.SaveAll {
		ss.Config.Params.SaveAll = false
		ss.Net.SaveParamsSnapshot(&ss.Config, ss.Config.Params.Good)
		os.Exit(0)
	}
}

func (ss *Sim) ConfigEnv() {
	// Can be called multiple times -- don't re-create
	var tst *env.FixedTable
	if len(ss.Envs) == 0 {
		tst = &env.FixedTable{}
	} else {
		tst = ss.Envs.ByMode(Test).(*env.FixedTable)
	}

	inputs := tensorfs.DirTable(ss.Root.Dir("Inputs/Digits"), nil)

	tst.Name = Test.String()
	tst.Config(table.NewView(inputs))
	tst.Sequential = true
	tst.Validate()

	tst.Init(0)

	// note: names must be in place when adding
	ss.Envs.Add(tst)
}

func (ss *Sim) ConfigNet(net *axon.Network) {
	net.SetMaxData(1)
	net.Context().SetMinusCycles(int32(ss.Config.Run.MinusCycles)).
		SetPlusCycles(int32(ss.Config.Run.PlusCycles)).Update()
	net.SetRandSeed(ss.RandSeeds[0]) // init new separate random seed, using run = 0

	inp := net.AddLayer2D("Input", axon.InputLayer, 7, 5)
	inp.Doc = "Input represents the visual appearance of different digits."
	recv := net.AddLayer2D("RecvNeuron", axon.SuperLayer, 1, 1)
	recv.Doc = fmt.Sprintf("RecvNeuron represents an individual neuron with synaptic weights tuned to detect the digit %d.", DigitWts)

	net.ConnectLayers(inp, recv, paths.NewFull(), axon.ForwardPath)
	recv.PlaceAbove(inp)

	net.Build()
	net.Defaults()
	ss.ApplyParams()
	ss.InitWeights(net)
}

// InitWeights initializes the weights, setting the receiving neuron's weights
// to the [DigitWts] digit pattern, which is what it then detects.
func (ss *Sim) InitWeights(net *axon.Network) {
	net.InitWeights()
	dpat := ss.Digits().Column("Input").RowTensor(DigitWts)
	recv := net.LayerByName("RecvNeuron")
	pt := recv.RecvPaths[0]
	for i := range dpat.Len() {
		pt.SetSynValue("Wt", i, 0, float32(dpat.Float1D(i)))
	}
}

func (ss *Sim) ApplyParams() {
	ss.Params.Script = ss.Config.Params.Script
	ss.Params.ApplyAll(ss.Net)

	recv := ss.Net.LayerByName("RecvNeuron")
	recv.Params.Acts.Gbar.L = ss.GbarL
}

////////  Init, utils

// Init restarts the run, and initializes everything, including network weights
// and resets the epoch log table
func (ss *Sim) Init() {
	ss.Loops.ResetCounters()
	ss.SetRunName()
	ss.InitRandSeed(0)
	ss.ApplyParams()
	ss.StatsInit()
	ss.NewRun()
	ss.NetUpdate.RecordSyns()
	ss.NetUpdate.Update(Test, Trial)
}

// InitRandSeed initializes the random seed based on current training run number
func (ss *Sim) InitRandSeed(run int) {
	ss.RandSeeds.Set(run, ss.Net.Rand)
}

// NetViewUpdater returns the NetViewUpdate for given mode.
func (ss *Sim) NetViewUpdater(mode enums.Enum) *axon.NetViewUpdate {
	return &ss.NetUpdate
}

// ConfigLoops configures the control loops: Training, Testing
func (ss *Sim) ConfigLoops() {
	ls := looper.NewStacks()

	ev := ss.Envs.ByMode(Test).(*env.FixedTable)
	ntrls := ev.Table.NumRows()
	cycles := ss.Config.Run.Cycles()

	ls.AddStack(Test, Trial).
		AddLevel(Epoch, 1).
		AddLevel(Trial, ntrls).
		AddLevel(Cycle, cycles)

	axon.LooperStandard(ls, ss.Net, ss.NetViewUpdater, Cycle, Trial, Train,
		func(mode enums.Enum) { ss.Net.ClearInputs() },
		func(mode enums.Enum) { ss.ApplyInputs(mode.(Modes)) },
	)
	ls.Stacks[Test].OnInit.Add("Init", ss.Init)

	ls.AddOnStartToAll("StatsStart", ss.StatsStart)
	ls.AddOnEndToAll("StatsStep", ss.StatsStep)

	if ss.Config.GUI {
		axon.LooperUpdateNetView(ls, Cycle, Trial, ss.NetViewUpdater)
		ls.Stacks[Test].OnInit.Add("GUI-Init", ss.GUI.UpdateWindow)
	}

	if ss.Config.Debug {
		mpi.Println(ls.DocString())
	}
	ss.Loops = ls
}

// ApplyInputs applies input patterns from given environment for given mode.
// Any other start-of-trial logic can also be put here.
func (ss *Sim) ApplyInputs(mode Modes) {
	net := ss.Net
	curModeDir := ss.Current.Dir(mode.String())
	ev := ss.Envs.ByMode(mode)
	lays := net.LayersByType(axon.InputLayer, axon.TargetLayer)
	net.InitExt()
	ev.Step()
	curModeDir.StringValue("TrialName", 1).SetString1D(ev.String(), 0)
	for _, lnm := range lays {
		ly := ss.Net.LayerByName(lnm)
		st := ev.State(ly.Name)
		if st != nil {
			ly.ApplyExt(uint32(0), st)
		}
	}
	net.ApplyExts()
}

// NewRun intializes a new Run level of the model.
func (ss *Sim) NewRun() {
	ctx := ss.Net.Context()
	ss.InitRandSeed(0)
	ss.Envs.ByMode(Test).Init(0)
	ctx.Reset()
	ss.InitWeights(ss.Net)
}

////////  Inputs

// OpenTable opens a [table.Table] from embedded content, storing
// the data in the given tensorfs directory.
func (ss *Sim) OpenTable(dir *tensorfs.Node, fsys fs.FS, fnm, name, docs string) (*table.Table, error) {
	dt := table.New()
	metadata.SetName(dt, name)
	metadata.SetDoc(dt, docs)
	err := dt.OpenFS(fsys, fnm, tensor.Tab)
	if errors.Log(err) != nil {
		return dt, err
	}
	tensorfs.DirFromTable(dir.Dir(name), dt)
	return dt, err
}

func (ss *Sim) OpenInputs() {
	dir := ss.Root.Dir("Inputs")
	ss.OpenTable(dir, embedfs, "digits.tsv", "Digits", "Digit testing patterns")
}

// Digits returns the table of digit input patterns.
func (ss *Sim) Digits() *table.Table {
	return tensorfs.DirTable(ss.Root.Dir("Inputs/Digits"), nil)
}

//////// Stats

// AddStatStd adds a standard stat compute function (defined in axon)
func (ss *Sim) AddStatStd(f func(mode enums.Enum, level enums.Enum, start bool)) {
	ss.StatFuncs = append(ss.StatFuncs, f)
}

// AddStat adds a custom stat compute function.
func (ss *Sim) AddStat(f func(mode Modes, level Levels, start bool)) {
	ss.AddStatStd(func(mode enums.Enum, level enums.Enum, start bool) {
		f(mode.(Modes), level.(Levels), start)
	})
}

// StatsStart is called by Looper at the start of given level, for each iteration.
// It needs to call RunStats Start at the next level down.
// e.g., each Epoch is the start of the full set of Trial Steps.
func (ss *Sim) StatsStart(lmd, ltm enums.Enum) {
	mode := lmd.(Modes)
	level := ltm.(Levels)
	if level < Trial {
		return
	}
	ss.RunStats(mode, level-1, axon.Start)
}

// StatsStep is called by Looper at each step of iteration,
// where it accumulates the stat results.
func (ss *Sim) StatsStep(lmd, ltm enums.Enum) {
	mode := lmd.(Modes)
	level := ltm.(Levels)
	ss.RunStats(mode, level, axon.Step)
	tensorfs.DirTable(axon.StatsNode(ss.Stats, mode, level), nil).WriteToLog()
}

// RunStats runs the StatFuncs for given mode, level and phase.
func (ss *Sim) RunStats(mode Modes, level Levels, start bool) {
	for _, sf := range ss.StatFuncs {
		sf(mode, level, start)
	}
	if !start && ss.GUI.Tabs != nil {
		nm := mode.String() + " " + level.String() + " Plot"
		ss.GUI.Tabs.AsLab().GoUpdatePlot(nm)
	}
}

// SetRunName sets the overall run name, used for naming output logs and weight files
// based on params extra sheets and tag, and starting run number (for distributed runs).
func (ss *Sim) SetRunName() string {
	runName := ss.Params.RunName(0)
	ss.Current.StringValue("RunName", 1).SetString1D(runName, 0)
	return runName
}

// RunName returns the overall run name, used for naming output logs and weight files
// based on params extra sheets and tag, and starting run number (for distributed runs).
func (ss *Sim) RunName() string {
	return ss.Current.StringValue("RunName", 1).String1D(0)
}

// StatsInit initializes all the stats by calling Start across all modes and levels.
func (ss *Sim) StatsInit() {
	for md, st := range ss.Loops.Stacks {
		mode := md.(Modes)
		for _, lev := range st.Order {
			level := lev.(Levels)
			ss.RunStats(mode, level, axon.Start)
		}
	}
	if ss.GUI.Tabs != nil {
		tbs := ss.GUI.Tabs.AsLab()
		_, idx := tbs.CurrentTab()
		tbs.PlotTensorFS(axon.StatsNode(ss.Stats, Test, Cycle))
		tbs.PlotTensorFS(axon.StatsNode(ss.Stats, Test, Trial))
		tbs.SelectTabIndex(idx)
	}
}

// statDocs are the docs for the receiving neuron stats.
var statDocs = map[string]string{
	"GeSyn":   "GeSyn is the excitatory AMPA synaptic conductance (net input) into the receiving neuron, which directly reflects the degree of match between the input pattern and the neuron's synaptic weights: GeSyn = (1 / Alpha) * Sum(Act * Wt) / N, where Alpha is the expected activity level of the sending layer.",
	"Ge":      "Ge is the total excitatory conductance into the receiving neuron, which adds the voltage-gated NMDA and VGCC channel contributions on top of the AMPA synaptic conductance in GeSyn. Because NMDA is voltage sensitive, this is not a linear function of the input overlap. At the Trial level this is the GeInt value integrated over the theta cycle.",
	"Act":     "Act is the activation (normalized firing rate) of the receiving neuron. At the Trial level this is the ActInt integrated value at the end of the trial.",
	"Vm":      "Vm is the membrane potential of the receiving neuron in mV, which integrates the excitatory, leak and inhibitory currents, and drives spiking when it gets above threshold.",
	"SpikeHz": "SpikeHz is the average firing rate of the receiving neuron in Hz (spikes per second) across the trial.",
}

// ConfigStats handles configures functions to do all stats computation
// in the tensorfs system.
func (ss *Sim) ConfigStats() {
	net := ss.Net
	ss.Stats = ss.Root.Dir("Stats")
	ss.Current = ss.Stats.Dir("Current")

	ss.SetRunName()

	// last arg(s) are levels to exclude
	ss.AddStatStd(axon.StatLoopCounters(ss.Stats, ss.Current, ss.Loops, net, Trial))
	ss.AddStatStd(axon.StatRunName(ss.Stats, ss.Current, ss.Loops, net, Trial))
	ss.AddStatStd(axon.StatTrialName(ss.Stats, ss.Current, ss.Loops, net, Trial))

	// The receiving neuron is the whole point of this sim, so we log its state
	// directly, at both the Cycle and Trial levels.
	statNames := []string{"GeSyn", "Ge", "Act", "Vm", "SpikeHz"}
	ss.AddStat(func(mode Modes, level Levels, start bool) {
		recv := ss.Net.LayerByName("RecvNeuron")
		ni := int(recv.NeurStIndex)
		di := 0
		for _, name := range statNames {
			modeDir := ss.Stats.Dir(mode.String())
			curModeDir := ss.Current.Dir(mode.String())
			levelDir := modeDir.Dir(level.String())
			tsr := levelDir.Float64(name)
			if start {
				tsr.SetNumRows(0)
				metadata.SetDoc(tsr, statDocs[name])
				plot.SetFirstStyler(tsr, func(s *plot.Style) {
					s.On = true
					switch name {
					case "GeSyn", "Act":
						s.Range.SetMin(0).SetMax(1)
					case "Ge":
						s.On = false
						s.Range.SetMin(0).SetMax(1)
					case "Vm":
						s.On = false
						s.Range.SetMin(-80).SetMax(-40)
						s.RightY = true
					case "SpikeHz":
						s.On = false
						s.Range.SetMin(0).SetMax(100)
						s.RightY = true
					}
				})
				continue
			}
			var stat float64
			switch level {
			case Cycle:
				switch name {
				case "GeSyn":
					stat = float64(axon.Neurons.Value(ni, di, int(axon.GeSyn)))
				case "Ge":
					stat = float64(axon.Neurons.Value(ni, di, int(axon.Ge)))
				case "Act":
					stat = float64(axon.Neurons.Value(ni, di, int(axon.Act)))
				case "Vm":
					stat = float64(axon.Neurons.Value(ni, di, int(axon.Vm)))
				case "SpikeHz":
					stat = 1000 * float64(axon.Neurons.Value(ni, di, int(axon.Spike)))
				}
			case Trial:
				subDir := modeDir.Dir(Cycle.String())
				switch name {
				case "GeSyn": // stable across the trial: take the final value
					stat = stats.StatFinal.Call(subDir.Value(name)).Float1D(0)
				case "Ge": // integrated over the theta cycle: much less noisy than Ge itself
					stat = float64(axon.Neurons.Value(ni, di, int(axon.GeInt)))
				case "Act":
					stat = float64(axon.Neurons.Value(ni, di, int(axon.ActInt)))
				default: // average over the cycles within this trial
					stat = stats.StatMean.Call(subDir.Value(name)).Float1D(0)
				}
			default:
				subDir := modeDir.Dir((level - 1).String())
				stat = stats.StatMean.Call(subDir.Value(name)).Float1D(0)
			}
			curModeDir.Float64(name, 1).SetFloat1D(stat, di)
			tsr.AppendRowFloat(stat)
		}
	})

	ss.AddStatStd(axon.StatPerTrialMSec(ss.Stats, Test, Trial))
}

// StatCounters returns counters string to show at bottom of netview.
func (ss *Sim) StatCounters(mode, level enums.Enum) string {
	counters := ss.Loops.Stacks[mode].CountersString()
	vu := ss.NetViewUpdater(mode)
	if vu == nil || vu.View == nil {
		return counters
	}
	di := vu.View.Di
	curModeDir := ss.Current.Dir(mode.String())
	if curModeDir.Node("TrialName") == nil {
		return counters
	}
	counters += fmt.Sprintf(" TrialName: %s", curModeDir.StringValue("TrialName").String1D(di))
	if curModeDir.Node("Ge") == nil {
		return counters
	}
	for _, name := range []string{"Ge", "Act"} {
		counters += fmt.Sprintf(" %s: %.4g", name, curModeDir.Float64(name).Float1D(di))
	}
	return counters
}

//////// GUI

// ConfigGUI configures the Cogent Core GUI interface for this simulation.
func (ss *Sim) ConfigGUI(b tree.Node) {
	ss.GUI.MakeBody(b, ss, ss.Root, ss.Config.Name, ss.Config.Title, ss.Config.Doc)
	ss.GUI.CycleUpdateInterval = 10
	ss.GUI.StopLevel = Trial

	nv := ss.GUI.AddNetView("Network")
	nv.Settings.MaxRecs = 2 * ss.Config.Run.Cycles()
	nv.Settings.Raster.Max = ss.Config.Run.Cycles()
	nv.SetNet(ss.Net)
	ss.NetUpdate.Config(nv, axon.Theta, ss.StatCounters)
	ss.GUI.OnStop = func(mode, level enums.Enum) {
		vu := ss.NetViewUpdater(mode)
		vu.UpdateWhenStopped(mode, level)
	}

	nv.SceneXYZ().Camera.Pose.Pos.Set(0, 1.5, 2.5)
	nv.SceneXYZ().Camera.LookAt(math32.Vec3(0, 0, 0), math32.Vec3(0, 1, 0))

	ss.StatsInit()
	ss.GUI.FinalizeGUI(false)
}

func (ss *Sim) MakeToolbar(p *tree.Plan) {
	ss.GUI.AddLooperCtrl(p, ss.Loops)

	tree.Add(p, func(w *core.Separator) {})
	ss.GUI.AddToolbarItem(p, egui.ToolbarItem{
		Label:   "Defaults",
		Icon:    icons.Update,
		Tooltip: "Restore initial default parameters.",
		Active:  egui.ActiveStopped,
		Func: func() {
			ss.Defaults()
			ss.Init()
			ss.GUI.UpdateWindow()
		},
	})
	ss.GUI.AddToolbarItem(p, egui.ToolbarItem{
		Label:   "README",
		Icon:    icons.FileMarkdown,
		Tooltip: "Opens your browser on the README file that contains instructions for how to run this model.",
		Active:  egui.ActiveAlways,
		Func: func() {
			core.TheApp.OpenURL(ss.Config.URL)
		},
	})
}

func (ss *Sim) RunNoGUI() {
	ss.Init()

	if ss.Config.Params.Note != "" {
		mpi.Printf("Note: %s\n", ss.Config.Params.Note)
	}

	runName := ss.SetRunName()
	netName := ss.Net.Name
	cfg := &ss.Config.Log
	axon.OpenLogFiles(ss.Loops, ss.Stats, netName, runName, [][]string{cfg.Save})

	ss.Loops.Run(Test)

	axon.CloseLogFiles(ss.Loops, ss.Stats)
	axon.GPURelease()
}
