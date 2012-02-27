// Berger - Parker diversity matrix
package div

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// The Berger-Parker index equals the maximum p[i] value in the dataset, i.e. the proportional abundance of the most abundant type.
// This corresponds to the weighted generalized mean of the p[i] values when q approaches infinity, and hence equals the inverse of true diversity of order infinity, 1/∞D.
func BergerParkerDiv(data *Matrix) *Vector {
	rows := data.R
	cols := data.C
	div := NewVector(rows)

	for i := 0; i < rows; i++ {
		tot := 0.0 // total number of all individuals in the sample
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				tot += x
			}
		}

		pmax := 0.0
		for j := 0; j < cols; j++ {
			p := data.Get(i, j) / tot
			pmax = math.Max(p, pmax)
		}
		div.Set(i, pmax)
	}
	return div
}
