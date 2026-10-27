// Copyright (c) 2025, The CCN Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sims has the simulations that the pages of the site embed, which
// js builds load only when the first one is shown, as Embed is in the
// Web.DeferredFunctions of core.toml.
package sims

import (
	"maps"
	"slices"

	"cogentcore.org/core/core"
	"cogentcore.org/core/system"
	"cogentcore.org/core/tree"
	"github.com/compcogneuro/web/sims/detector"
	"github.com/compcogneuro/web/sims/stability"
	"github.com/compcogneuro/web/sims/urakubo"
	"github.com/emer/axon/v2/sims/bgdorsal"
	"github.com/emer/axon/v2/sims/bgventral"
	"github.com/emer/axon/v2/sims/deepfsa"
	"github.com/emer/axon/v2/sims/inhib"
	"github.com/emer/axon/v2/sims/neuron"
)

// Embed embeds the simulation of the given name in the given parent,
// returning false if there is no simulation of that name.
func Embed(parent tree.Node, name string) bool {
	em, ok := sims[name]
	if ok {
		em(parent)
	}
	return ok
}

// Names returns the names of the simulations.
func Names() []string {
	return slices.Sorted(maps.Keys(sims))
}

func useGPU() bool {
	return core.TheApp.Platform() != system.Web
}

func ndata(nd int) int {
	if core.TheApp.Platform() == system.Web {
		return 1
	}
	return nd
}

// sims are the functions that embed each simulation, by name.
var sims = map[string]func(tree.Node){
	"bgdorsal": func(b tree.Node) {
		bgdorsal.Embed(b, func(ss *bgdorsal.Sim, cfg *bgdorsal.Config) {
			cfg.GPU = useGPU()
			cfg.Run.NData = ndata(cfg.Run.NData)
		})
	},
	"bgventral": func(b tree.Node) {
		bgventral.Embed(b, func(ss *bgventral.Sim, cfg *bgventral.Config) {
			cfg.GPU = useGPU()
			cfg.Run.NData = ndata(cfg.Run.NData)
		})
	},
	"deepfsa": func(b tree.Node) {
		deepfsa.Embed(b, func(ss *deepfsa.Sim, cfg *deepfsa.Config) {
			cfg.GPU = useGPU()
			cfg.Run.NData = ndata(cfg.Run.NData)
		})
	},
	"detector": func(b tree.Node) {
		detector.Embed(b, func(ss *detector.Sim, cfg *detector.Config) {
			cfg.GPU = useGPU()
		})
	},
	"inhib": func(b tree.Node) {
		inhib.Embed(b, func(ss *inhib.Sim, cfg *inhib.Config) {
			cfg.GPU = useGPU()
		})
	},
	"neuron": func(b tree.Node) {
		neuron.Embed(b, func(ss *neuron.Sim, cfg *neuron.Config) {
			cfg.GPU = useGPU()
		})
	},
	"stability": func(b tree.Node) {
		stability.Embed(b, func(ss *stability.Sim, cfg *stability.Config) {
			cfg.GPU = useGPU()
		})
	},
	"urakubo": func(b tree.Node) {
		urakubo.Embed(b, func(ss *urakubo.Sim, cfg *urakubo.Config) {
			cfg.GPU = false
		})
	},
}
