// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"math"

	"github.com/cpmech/gosl/gm/msh"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
	"github.com/cpmech/gosl/plt"
	"github.com/cpmech/gosl/utl"
)

func main() {

	// run profile
	defer utl.ProfCPU("/tmp/gosl", "cpu.integ", false)()

	// integrand function for second moment of inertia about x-axis: Ix
	fcnIx := func(x la.Vector) float64 {
		return x[1] * x[1]
	}

	// constants
	r, R := 1.0, 3.0
	nr, na := 10, 35
	anaIx := math.Pi * (math.Pow(R, 4) - math.Pow(r, 4)) / 4.0

	// generate mesh
	ctype := msh.TypeQua17
	mesh := msh.GenRing2d(ctype, nr, na, r, R, 2.0*math.Pi)

	// allocate cell integrator with default integration points
	o := msh.NewMeshIntegrator(mesh, 1)

	// compute Ix
	Ix := o.IntegrateSv(0, fcnIx)

	// compare with analytical solution
	typekey := msh.TypeIndexToKey[ctype]
	io.Pf("%s : Ix = %v  error = %v\n", typekey, Ix, math.Abs(Ix-anaIx))

	// plot mesh
	if true {
		plt.Reset(true, &plt.A{WidthPt: 400, Dpi: 150})
		args := msh.NewArgs()
		args.WithEdges = true
		args.WithVerts = true
		args.WithCells = false
		mesh.Draw(args)
		plt.HideAllBorders()
		plt.Save("/tmp/gosl/gm", io.Sf("integ04-%s", typekey))
	}
}
