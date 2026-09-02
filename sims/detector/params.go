// Copyright (c) 2026, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package detector

import "github.com/emer/axon/v2/axon"

// LayerParams sets the minimal non-default params.
// Base is always applied, and others can be optionally selected to apply on top of that.
var LayerParams = axon.LayerSheets{
	"Base": {
		{Sel: "Layer", Doc: "no inhibition: there is only one receiving neuron here, and the point is to see the raw excitatory drive",
			Set: func(ly *axon.LayerParams) {
				ly.Inhib.Layer.On.SetBool(false)
				ly.Inhib.ActAvg.AdaptGi.SetBool(false)
			}},
		{Sel: "#Input", Doc: "set expected activity of input layer: key for normalizing the net input. 17 of the 35 units are active in the digit 8 pattern that the weights encode.",
			Set: func(ly *axon.LayerParams) {
				ly.Inhib.ActAvg.Nominal = 0.4857 // 17 / 35
			}},
	},
}

// PathParams sets the minimal non-default params.
// Base is always applied, and others can be optionally selected to apply on top of that.
var PathParams = axon.PathSheets{
	"Base": {
		{Sel: "Path", Doc: "no learning: weights are set directly to the digit 8 pattern",
			Set: func(pt *axon.PathParams) {
				pt.Learn.Learn.SetBool(false)
				// Because the input layer is clamped, its spike-driven conductance
				// integrates to 1.4874x the analytical Sum(Act*Wt)/N net input value.
				// Dividing that factor out here makes GeSyn exactly match the
				// analytical net input equation given in the README, so that a
				// perfectly matching input pattern gives GeSyn = 1.
				pt.PathScale.Abs = 1 / 1.4874
			}},
	},
}
