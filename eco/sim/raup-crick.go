// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Raup - Crick distance and similarity. 
// Raup & Crick (1979). 

import (
	. "code.google.com/p/go-eco/eco/aux"
	"code.google.com/p/go-fn/fn"
	"code.google.com/p/probab/dst"
	"math"
)

func probK(a, b, n, k float64) float64 {
	logNum1 := fn.LnFactBig(b) - (fn.LnFactBig(b-k) + fn.LnFactBig(k))
	logNum2 := fn.LnFactBig(n-b) - (fn.LnFactBig(n-b-a+k) + fn.LnFactBig(a-k))
	logDen := fn.LnFactBig(n) - (fn.LnFactBig(n-a) + fn.LnFactBig(a))
	return math.Exp(logNum1 + logNum2 - logDen)
}

// RaupCrickBool_D returns a Raup - Crick distance matrix for boolean data. 
// Algorithm from R:vegan. 
func RaupCrickBool_D(data *Matrix) *Matrix {
	// Raup - Crick distance is a probabilistic index based on presence/absence data. It is defined as 1 - prob(j), 
	// or based on the probability of observing at least j species in shared in compared communities. 
	// Legendre & Legendre (1998) suggest using simulations to assess the probability, but the current function uses analytic result 
	// from hypergeometric distribution instead. This probability (and the index) is dependent on the number of species missing in both sites, 
	// and adding all-zero species to the data or removing missing species from the data will influence the index. 
	// The probability (and the index) may be almost zero or almost one for a wide range of parameter values. 
	// The index is nonmetric: two communities with no shared species may have a dissimilarity slightly below one, 
	// and two identical communities may have dissimilarity slightly above zero. 
	// Compared to other metrics for p/a data, Raup-Crick seems to be very robust for small samples.
	// Algorithm from R:vegan
	// phyper(k, m, size-m, n) == Hypergeometric_CDF_At(size, m, n, k)

	var (
		v                          float64
		aaa, bbb, jjj, t1, t2, sim int64
	)

	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)
	WarnIfNotBool(data)
	WarnIfEmptySpecies(data)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sim = 0
			t1 = 0
			t2 = 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)

				if x > 0.0 && y > 0.0 {
					sim++
				}
				if x > 0.0 {
					t1++
				}
				if y > 0.0 {
					t2++
				}
			}

			jjj = sim - 1
			if t1 < t2 {
				aaa = t1
				bbb = t2

			} else {
				aaa = t2
				bbb = t1
			}
			//	v = 1 - phyper(jjj, aaa, float64(count) - aaa, bbb, 1, 0);

			//fmt.Println("hyper: ", cols, aaa, bbb, jjj)
			v = 1.0 - dst.HypergeometricCDFAt(int64(cols), aaa, bbb, jjj)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// RaupCrickBool1_S returns a Raup - Crick similarity matrix #1 for boolean data. 
// Raup & Crick (1979): 1217, eq. 4
// This is the naive version of their similarity index;
// for final version, use the algorithm described on page 1219
func RaupCrickBool1_S(data *Matrix) *Matrix {
	var a, b, n float64

	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)
	WarnIfNotBool(data)

	n = float64(cols)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a = 0
			b = 0
			common := 0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				if x > 0.0 && y > 0.0 {
					common++
				}

				if x > 0.0 {
					a++
				}
				if y > 0.0 {
					b++
				}
			}

			p := 0.0
			for k := 0; k < common; k++ {
				p += probK(a, b, n, float64(k))
			}

			out.Set(i, j, p)
			out.Set(j, i, p)
		}
	}
	return out
}

// RaupCrickBool2_S returns a Raup - Crick similarity matrix #2 for boolean data. 
// Raup & Crick (1979): 1219. 
// This is the final version of their similarity index. 
func RaupCrickBool2_S(data *Matrix, p []float64) *Matrix {
	const iter int = 1e3 // number of iterations, adjust according to speed

	rows := data.R
	cols := data.C
	out := NewMatrix(rows, rows)
	WarnIfNotBool(data)

	a := make([]int, cols)
	b := make([]int, cols)
	k := make([]int, cols)

	// if p == nil, estimate it
	if p == nil {
		p = make([]float64, cols)

		//get grand total
		gTot := 0.0
		for i := 0; i < cols; i++ {
			for j := i; j < rows; j++ {
				gTot += data.Get(i, j)
			}
		}

		for i := 0; i < cols; i++ {
			sum := 0.0 // this species' total over all samples
			for j := i; j < rows; j++ {
				sum += data.Get(i, j)
			}
			p[i] = sum / gTot
		}
	}

	// for all pairs of samples
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			// count species in rows i and j
			aCount := 0
			bCount := 0
			k_obs := 0
			for l := 0; l < cols; l++ {
				if data.Get(i, l) > 0 {
					aCount++
				}
				if data.Get(j, l) > 0 {
					bCount++
				}
				if data.Get(i, l) > 0 && data.Get(j, l) > 0 {
					k_obs++
				}
			}
			for l := 0; l < cols; l++ {
				k[l] = 0
			}

			// accumulate counts for k_exp
			for l := 0; l < iter; l++ {

				// create assemblage A
				nSp := 0
				for m := 0; m < cols; m++ {
					a[m] = 0
				}
			L1:
				for {
					// draw from categorical ditribution
					cat := dst.ChoiceNext(p)
					// add the species to assemblage, if new
					if a[cat] == 0 {
						a[cat] = 1
						nSp++
					}
					// if number of species == aCount, break to L1
					if nSp == aCount {
						break L1
					}
				}

				// create assemblage B
				nSp = 0
				for m := 0; m < cols; m++ {
					b[m] = 0
				}
			L2:
				for {
					// draw from categorical ditribution
					cat := dst.ChoiceNext(p)
					// add the species to assemblage, if new
					if b[cat] == 0 {
						b[cat] = 1
						nSp++
					}
					// if number of species == bCount, break to L2
					if nSp == bCount {
						break L2
					}
				}

				k_exp := 0
				// count species in common for A and B
				for m := 0; m < cols; m++ {
					if a[m] > 0 && b[m] > 0 {
						k_exp++
					}
				}
				k[k_exp]++ // add it to histogram
			} // end of iterations

			// sim = CDF(k_obs - 1) + PDF(k_obs)/2
			prob := 0.0
			for l := 0; l < k_obs; l++ {
				prob += float64(k[l]) / float64(iter)
			}
			prob += float64(k[k_obs]) / (2 * float64(iter))

			out.Set(i, j, prob)
			out.Set(j, i, prob)

		}
	}
	return out
}
