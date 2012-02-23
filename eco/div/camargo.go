// Camargo E equitability (evenness)
// Camargo, J. A. 1992b. New diversity index for assessing structural alterations in aquatic communities. - Bull. Environ. Contam. Toxicol. 48: 428-434.
// Camargo, J. A. 1993a. Must dominance increase with the number of subordinate species in competitive interactions? - J. Theor. Biol. 161: 537-542.

package div

import (
	. "go-eco.googlecode.com/hg/eco"
)

// Camargo E equitability (evenness)
// Camargo 1992b, 1993a, 1995
func Camargo_E(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		sumX := 0.0 // total number of all individuals in the sample
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			sumX += x
		}

		// recalculate counts to proportions
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			data.Set(i, j, x/sumX)
		}
		
		v:= 0.0
		for j := 0; j < cols; j++ {
			for k := j+1; k < cols; k++ {
				pj := data.Get(i, j)
				pk := data.Get(i, k)
				v += (pj-pk)/sumX
			}
		}
		out.Set(i, v)
	}
	return out
}


