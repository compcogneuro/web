// Copyright (c) 2026, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package detector

import (
	"math"
	"testing"

	"cogentcore.org/lab/table"
	"cogentcore.org/lab/tensorfs"
	"github.com/emer/axon/v2/axon"
	"github.com/emer/emergent/v2/egui"
)

// overlaps is the number of active units in each digit pattern that also have
// a weight of 1 in the [DigitWts] pattern that the receiving neuron detects.
var overlaps = []float64{6, 6, 12, 13, 5, 14, 12, 6, 17, 12}

// runTest runs the full Test stack with the given leak conductance,
// returning the Test Trial stats table.
func runTest(gbarL float32) *table.Table {
	cfgC, _ := egui.NewConfig[Config]()
	cfgC.GUI = false
	cfgC.GPU = false
	ss := &Sim{}
	ss.SetConfig(cfgC)
	ss.ConfigSim()
	ss.GbarL = gbarL
	ss.Init()
	ss.Loops.Run(Test)
	return tensorfs.DirTable(axon.StatsNode(ss.Stats, Test, Trial), nil)
}

// TestGeSyn tests that the excitatory net input into the receiving neuron
// matches the analytical net input equation given in the README:
// GeSyn = (1 / Alpha) * Sum(Act * Wt) / N, which for these binary patterns
// and weights reduces to the overlap between the input pattern and the
// weights, divided by the number of active units in the weight pattern.
func TestGeSyn(t *testing.T) {
	dt := runTest(240)
	if dt.NumRows() != len(overlaps) {
		t.Fatalf("expected %d trials, got %d", len(overlaps), dt.NumRows())
	}
	ge := dt.Column("GeSyn")
	for r, ov := range overlaps {
		want := ov / overlaps[DigitWts]
		got := ge.FloatRow(r, 0)
		if math.Abs(got-want) > 0.005 {
			t.Errorf("digit %d: GeSyn = %g, expected %g (overlap %g)", r, got, want, ov)
		}
	}
}

// TestDetector tests that at the default leak level, only the digit that the
// weights encode activates the receiving neuron, and that lowering the leak
// makes it respond to the next-closest patterns as well.
func TestDetector(t *testing.T) {
	tests := []struct {
		gbarL  float32
		active []int
	}{
		{240, []int{8}},                // only the perfect match (overlap 17)
		{220, []int{5, 8}},             // + overlap 14
		{200, []int{3, 5, 8}},          // + overlap 13
		{180, []int{2, 3, 5, 6, 8, 9}}, // + all three with overlap 12
		{300, nil},                     // too much leak: nothing gets through
	}
	for _, tst := range tests {
		dt := runTest(tst.gbarL)
		act := dt.Column("Act")
		var got []int
		for r := range dt.NumRows() {
			if act.FloatRow(r, 0) > 0.1 {
				got = append(got, r)
			}
		}
		if len(got) != len(tst.active) {
			t.Errorf("GbarL %g: active digits %v, expected %v", tst.gbarL, got, tst.active)
			continue
		}
		for i, d := range tst.active {
			if got[i] != d {
				t.Errorf("GbarL %g: active digits %v, expected %v", tst.gbarL, got, tst.active)
				break
			}
		}
	}
}
