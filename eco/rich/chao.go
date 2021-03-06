// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package rich

// Chao species number estimators. 
// Robert K. Colwell, Anne Chao, Nicholas J. Gotelli, Shang-Yi Lin, Chang Xuan Mao, Robin L. Chazdon,  and John T. Longino 2012: Models and estimators linking individual-based and sample-based rarefaction, extrapolation and comparison of assemblages J Plant Ecol (2012) 5(1): 3-21 doi:10.1093/jpe/rtr044.
// These nonparametic estimators of species richness are minimum estimators: their computed values should be viewed as lower bounds of total species numbers, given the information in a sample or sample set.

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Estimators ...

// ChaoS returns a vector of the Chao species number estimator for abundance data, classical formula. 
// Chao (1984, 1987). 
func ChaoS(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	aux.ToCounts(data)

	for i := 0; i < rows; i++ {
		s0 := 0.0
		s1 := 0.0
		s2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				s0++
				if x == 1 {
					s1++
				} else if x == 2 {
					s2++
				}
			}

		}
		v := s0 + s1*s1/(s2*2)
		out.Set(i, v)
	}
	return out
}

// ChaoAutoS returns a vector of  the Chao species estimator for abundance data, auto-corrected for bias when classical returns NaN. 
// Chao (1984, 1987). 
func ChaoAutoS(data *aux.Matrix) *aux.Vector {
	var v float64
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	aux.ToCounts(data)

	for i := 0; i < rows; i++ {
		s0 := 0.0
		s1 := 0.0
		s2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				s0++
				if x == 1 {
					s1++
				} else if x == 2 {
					s2++
				}
			}

		}
		if (s1-s2)*(s1-s2) == (s1+s2)*(s1+s2) {
			v = s0 + s1*(s1-1)/((s2+1)*2)
		} else {
			v = s0 + s1*s1/(s2*2)
		}
		out.Set(i, v)
	}
	return out
}

// ChaoCorrS  returns a vector of  the bias-corrected Chao species estimator for abundance data. 
// Chao (1984, 1987). 
func ChaoCorrS(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	aux.ToCounts(data)

	for i := 0; i < rows; i++ {
		s0 := 0.0
		s1 := 0.0
		s2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				s0++
				if x == 1 {
					s1++
				} else if x == 2 {
					s2++
				}
			}

		}
		v := s0 + s1*(s1-1)/((s2+1)*2) // return bias-corrected estimate
		out.Set(i, v)
	}
	return out
}

// ChaoBoolS  returns a vector of  the Chao species estimator for boolean (presence-absence) data, classical version. 
// Chao 1984, 1987
func ChaoBoolS(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	aux.ToBool(data)

	for i := 0; i < rows; i++ {
		s0 := 0.0
		s1 := 0.0
		s2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				s0++
				if x == 1 {
					s1++
				} else if x == 2 {
					s2++
				}
			}

		}
		v := s0 + s1*s1/(s2*2)
		out.Set(i, v)
	}
	return out
}

// ChaoAutoBoolS  returns a vector of the Chao species estimator for boolean (presence-absence) data, auto-corrected for bias when classical returns NaN. 
// Chao (1984, 1987). 
func ChaoAutoBoolS(data *aux.Matrix) *aux.Vector {
	var v float64
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	aux.ToBool(data)

	for i := 0; i < rows; i++ {
		s0 := 0.0
		s1 := 0.0
		s2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				s0++
				if x == 1 {
					s1++
				} else if x == 2 {
					s2++
				}
			}

		}
		if (s1-s2)*(s1-s2) == (s1+s2)*(s1+s2) {
			v = s0 + s1*(s1-1)/((s2+1)*2)
		} else {
			v = s0 + s1*s1/(s2*2)
		}
		out.Set(i, v)
	}
	return out
}

// ChaoCorrBoolS returns a vector of the bias-corrected  Chao species estimator for boolean (presence-absence) data. 
// Chao (1984, 1987). 
func ChaoCorrBoolS(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	aux.ToBool(data)

	for i := 0; i < rows; i++ {
		s0 := 0.0
		s1 := 0.0
		s2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				s0++
				if x == 1 {
					s1++
				} else if x == 2 {
					s2++
				}
			}

		}
		v := s0 + s1*(s1-1)/((s2+1)*2)
		out.Set(i, v)
	}
	return out
}

// ... and their variances

// ChaoVar returns a vector of variances of the Chao species estimator for abundance (count) data, classical formula. 
// Chao (1984, 1987). 
func ChaoVar(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	aux.ToCounts(data)

	for i := 0; i < rows; i++ {
		f1 := 0.0
		f2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				if x == 1 {
					f1++
				} else if x == 2 {
					f2++
				}
			}

		}
		fRat := f1 / f2
		fRat2 := fRat * fRat
		fRat3 := fRat * fRat * fRat
		fRat4 := fRat * fRat * fRat * fRat
		v := f2 * (0.5*fRat2 + fRat3 + 0.25*fRat4)
		out.Set(i, v)
	}
	return out
}

// ChaoCorrVar returns a vector of variances  of the bias-corrected Chao species estimator for abundance (count) data. 
// Chao (1984, 1987). 
func ChaoCorrVar(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	aux.ToCounts(data)

	for i := 0; i < rows; i++ {
		f1 := 0.0
		f2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				if x == 1 {
					f1++
				} else if x == 2 {
					f2++
				}
			}

		}
		a := f1 * (f1 - 1) / (2 * (f2 + 1))
		b := f1 * (2*f1 - 1) * (2*f1 - 1) / (4 * (f2 + 1) * (f2 + 1))
		c := f1 * f1 * f2 * (f1 - 1) * (f1 - 1) / (4 * (f2 + 1) * (f2 + 1) * (f2 + 1) * (f2 + 1))
		v := a + b + c
		out.Set(i, v)
	}
	return out
}

// ChaoBoolVar returns a vector of variances of the Chao species estimator for boolean (presence-absence) data, classical version. 
// Chao (1984, 1987). 
func ChaoBoolVar(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	aux.ToBool(data)

	for i := 0; i < rows; i++ {
		q1 := 0.0
		q2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				if x == 1 {
					q1++
				} else if x == 2 {
					q2++
				}
			}

		}
		qRat := q1 / q2
		qRat2 := qRat * qRat
		qRat3 := qRat * qRat * qRat
		qRat4 := qRat * qRat * qRat * qRat
		v := q2 * (0.5*qRat2 + qRat3 + 0.25*qRat4)
		out.Set(i, v)
	}
	return out
}

// ChaoBoolCorrVar returns a vector of variances of the Chao species estimator for boolean (presence-absence) data, bias-corrected version
// Chao (1984, 1987). 
func ChaoBoolCorrVar(data *aux.Matrix) *aux.Vector {
	rows := data.R
	cols := data.C
	out := aux.NewVector(rows)

	aux.ToBool(data)

	for i := 0; i < rows; i++ {
		q1 := 0.0
		q2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				if x == 1 {
					q1++
				} else if x == 2 {
					q2++
				}
			}

		}
		m := float64(rows)
		m1 := (m - 1) / m
		a := m1 * q1 * (q1 - 1) / (2 * (q2 + 1))
		b := m1 * m1 * q1 * (2*q1 - 1) * (2*q1 - 1) / (4 * (q2 + 1) * (q2 + 1))
		c := m1 * m1 * q1 * q1 * q2 * (q1 - 1) * (q1 - 1) / (4 * (q2 + 1) * (q2 + 1) * (q2 + 1) * (q2 + 1))
		v := a + b + c
		out.Set(i, v)
	}
	return out
}

// ChaoCI returns the 95% confidence interval of the Chao species estimator. 
// Chao (1984, 1987). 
func ChaoCI(sObs, chao, variance *aux.Vector) (lo, hi *aux.Vector) {
	cols := sObs.L
	if chao.L != cols || variance.L != cols {
		panic("bad data: unequal lengths")
	}
	for i := 0; i < cols; i++ {
		s := sObs.Get(i)
		c := chao.Get(i)
		v := variance.Get(i)
		t := c - s
		k := math.Exp(1.96 * math.Sqrt(math.Log(1+(v/(t*t)))))
		lo.Set(i, s+t/k)
		hi.Set(i, s+t*k)
	}
	return
}
