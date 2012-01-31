// Ruggiero similarity matrix
// Ruggiero et al. (1998), Koleff et al. (2003)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// Ruggiero similarity matrix
func RuggieroBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, _, c, _ = getABCD(data, i, j)
			v := a / (a + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}