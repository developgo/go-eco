// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package div

// Shannon diversity. 
import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// ShannonDiv returns vector of Shannon diversities. 
func ShannonDiv(data *aux.Matrix, base byte, corr bool) *aux.Vector {
	var log float64
	rows := data.R
	cols := data.C
	div := aux.NewVector(rows)

	for i := 0; i < rows; i++ {
		h := 0.0
		tot := 0.0 // total number of all individuals in the sample
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			tot += x
		}

		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			switch {
			case base == '2':
				log = math.Log2(x / tot)
			case base == 'd':
				log = math.Log10(x / tot)
			case base == 'e':
				log = math.Log(x / tot)
			default:
				log = math.Log(x / tot)
			}

			//    H += (x[i]/nn)*log(x[i]/nn);
			h -= (x / tot) * log
		}
		// correction term
		if corr {
			h -= float64(cols-1) / (2 * tot)
		}
		div.Set(i, h)
	}
	return div
}

// ShannonMax returns ...
func ShannonMax(spec int64, base byte) float64 {
	var x float64
	s := float64(spec)
	switch {
	case base == '2':
		x = math.Log2(s)
	case base == 'd':
		x = math.Log10(s)
	case base == 'e':
		x = math.Log(s)
	default:
		x = math.Log(s)
	}
	return x
}

// ShannonMin returns ...
func ShannonMin(tot, spec int64, base byte) float64 {
	//tot	total number of all individuals,  tot = sum(x[j]) 
	var x float64
	t := float64(tot)
	s := float64(spec)

	switch {
	case base == '2':
		x = math.Log2(t) - ((t - s + 1) * math.Log2(t-s+1) / t)
	case base == 'd':
		x = math.Log10(t) - ((t - s + 1) * math.Log10(t-s+1) / t)
	case base == 'e':
		x = math.Log(t) - ((t - s + 1) * math.Log(t-s+1) / t)
	default:
		s = math.Log(t) - ((t - s + 1) * math.Log(t-s+1) / t)
	}
	return x
}
