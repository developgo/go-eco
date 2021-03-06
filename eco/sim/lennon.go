// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Lennon dissimilarity matrix
// Lennon et al. (2001), Koleff et al. (2003)

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Lennon1Bool_D returns a Lennon dissimilarity matrix #1 for boolean data. 
func Lennon1Bool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := (2 * math.Abs(b-c)) / (2*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Lennon2Bool_D returns a Lennon dissimilarity matrix #2 for boolean data. 
func Lennon2Bool_D(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := 1 - (math.Log((2*a+b+c)/(a+b+c)) / math.Log(2))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
