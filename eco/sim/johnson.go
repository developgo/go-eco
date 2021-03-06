// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Johnson similarity matrix
// Johnson (1971), Johnson (1967)

import (
	"code.google.com/p/go-eco/eco/aux"
)

// Johnson1Bool_S returns a Johnson similarity matrix #1 for boolean data. 
// Johnson (1971). 
func Johnson1Bool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, _, _ = aux.GetABCD(data, i, j)
			v := a / (2 * b)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Johnson2Bool_S returns a Johnson similarity matrix #2 for boolean data. 
// Johnson (1967). 
func Johnson2Bool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (a / (a + b)) + (a / (a + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
