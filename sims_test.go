// Copyright (c) 2026, The CCN Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"slices"
	"testing"

	"github.com/compcogneuro/web/sims"
	"github.com/stretchr/testify/assert"
)

// TestSimNames checks that simNames, which the app has without loading the
// sims package, are the names of its simulations.
func TestSimNames(t *testing.T) {
	assert.Equal(t, sims.Names(), slices.Sorted(slices.Values(simNames)))
}
