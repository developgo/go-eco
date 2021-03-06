// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Steinhaus similarity
// Motyka (1947). 
// Legendre & Legendre (1998): 265, eq. 7.24 (S17 index). 
// For count or interval data. 

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Steinhaus similarity matrix
// Steinhaus_S returns a Steinhaus similarity matrix for floating-point data. 
// Legendre & Legendre (1998): 265, eq. 7.24 (S17 index). 
func Steinhaus_S(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumX := 0.0
			sumY := 0.0
			sumMin := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumX += x
				sumY += y
				sumMin += math.Min(x, y)
			}
			v := 2 * sumMin / (sumX + sumY)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
