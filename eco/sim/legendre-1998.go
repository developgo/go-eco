// Similarity matrices from Legendre & Legendre (1998), wrappers for named functions

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

// S1 matrix
// Legendre & Legendre (1998): 255, eq. 7.1 (S1 index)
func S1(data *DenseMatrix) *DenseMatrix {
	return SimpleMatchingBool_S(data)
}

// S2 matrix
// Legendre & Legendre (1998): 255, eq. 7.2 (S2 index)
func S2(data *DenseMatrix) *DenseMatrix {
	return RogersTanimotoBool_S(data)
}

// S3 matrix
// Legendre & Legendre (1998): 255, eq. 7.3 (S3 index)
func S3(data *DenseMatrix) *DenseMatrix {
	return SokalSneath1Bool_S(data)
}

// S4 matrix
// Legendre & Legendre 1998: 255, eq. 7.4 (S4 index)
func S4(data *DenseMatrix) *DenseMatrix {
	return SokalSneath2Bool_S(data)
}

// S5 matrix
// Legendre & Legendre 1998: 255, eq. 7.5 (S5 index)
func S5(data *DenseMatrix) *DenseMatrix {
	return SokalSneath3Bool_S(data)
}

// S6 matrix
// Legendre & Legendre 1998: 255, eq. 7.6 (S6 index)
func S6(data *DenseMatrix) *DenseMatrix {
	return SokalSneath4Bool_S(data)
}

// S7 matrix
// Legendre & Legendre 1998: 256, eq. 7.10 (S7 index)
func S7(data *DenseMatrix) *DenseMatrix {
	return JaccardBool_S(data)
}

// S8 matrix
// Legendre & Legendre (1998): 256, eq. 7.11 (S8 index)
func S8(data *DenseMatrix) *DenseMatrix {
	return SorensenBool_S(data)
}

// S9 similarity matrix, for boolean data
// Legendre & Legendre (1998): 257, eq. 7.12 (S9 index)
func S9(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	out := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			v := 3 * a / (3*a + b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}



// S10 matrix
// Legendre & Legendre 1998: 255, eq. 7.13 (S10 index)
func S10(data *DenseMatrix) *DenseMatrix {
	return SokalSneath5Bool_S(data)
}

// S11 matrix
// Legendre & Legendre (1998): 257, eq. 7.14 (S11 index)
func S11(data *DenseMatrix) *DenseMatrix {
	return RusselRaoBool_S(data)
}

// S12 matrix
// Legendre & Legendre (1998): 257, eq. 7.15 (S12 index)
func S12(data *DenseMatrix) *DenseMatrix {
	return Kulczynski1Bool_S(data)
}

// S13 matrix
// Legendre & Legendre 1998: 255, eq. 7.16 (S13 index)
func S13(data *DenseMatrix) *DenseMatrix {
	return SokalSneath9Bool_S(data)
}

// S14 matrix
// Legendre & Legendre (1998): 258, eq. 7.17 (S14 index)
func S14(data *DenseMatrix) *DenseMatrix {
	return OchiaiBool_S(data)
}

// S15 matrix == Gower
// Legendre & Legendre (1998): 259, eq. 7.20 (S15 index)
// TO BE IMPLEMENTED

// S16 matrix == Estabrook & Rogers
// Legendre & Legendre (1998): 261, eq. 7.21 (S16 index)
// TO BE IMPLEMENTED









// S26 matrix
// Legendre & Legendre (1998): 258, eq. 7.18 (S26 index)
func S26(data *DenseMatrix) *DenseMatrix {
	return FaithBool_S(data)
}
