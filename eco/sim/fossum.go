// Fossum similarity matrix
// Holliday et al. (2002), Ellis et al. (1993)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Fossum similarity matrix
func FossumBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = getABCD(data, i, j)
			v := ((a + b + c + d) * (-1 * ((a / 2) * (a / 2)))) / ((a + b) * (a + c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}