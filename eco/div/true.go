// True diversity matrix
package div

import (
	. "code.google.com/p/go-eco/eco"
	"math"
)

// True diversity, or the effective number of types, refers to the number of equally-abundant types needed for the average proportional abundance of the types to equal that observed in the dataset of interest
// (where all types may not be equally abundant). The true diversity in a dataset is calculated by first taking the weighted average of the proportional abundances of the types in the dataset,
// and then taking the inverse of this. The equation is:[3][4]
//    {}^q\!D={1 \over \sqrt[q-1]{{\sum_{i=1}^R p_i p_i^{q-1}}}}
// The denominator equals average proportional abundance of the types in the dataset as calculated with the weighted generalized mean with exponent q - 1. 
// In the equation, R is richness or the total number of types in the dataset, and the proportional abundance of the ith type is pi. The proportional abundances themselves are used as weights.
// The value of q defines which kind of mean is used: q = 0 corresponds to the harmonic mean, q = 1 to the geometric mean and q = 2 to the arithmetic mean. Increasing the value of q increases the weight given to the most abundant species, and hence causes the generalized mean to increase. As q approaches infinity, the generalized mean approaches the maximum pi value. In theory, q could also take negative values, but this would give rare species more weight than abundant ones, and hence cause the effective number of species to exceed the actual number of species. Negative values of q are therefore not used when the interest is in diversity.[3][4]
// The above equation is often written in the equivalent form:[1][2]
//    {}^q\!D=\left ( {\sum_{i=1}^R p_i^q} \right )^{1/(1-q)}
// The term inside the parentheses is called the basic sum. Some popular diversity indices correspond to the basic sum as calculated with different values of q.[2]
func True(data *Matrix, q float64) *Vector {
	rows := data.R
	cols := data.C
	div := NewVector(cols)

	for i := 0; i < rows; i++ {
		sum := 0.0
		tot := 0.0 // total number of all individuals in the sample
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0 {
				tot += x
			}
		}

		for j := 0; j < cols; j++ {
			p := data.Get(i, j) / tot
			sum += p * math.Pow(p, q-1)
		}
		div.Set(i, 1/math.Pow(sum, 1/(q-1)))
	}
	return div
}

// Not used now, rewrite for q == 0, 1, 2
// q == 0 --> harmonic
// q == 1 --> geometric
// q == 2 --> arithmetic

// Harmonic mean
func harmonicMean(data *Vector) float64 {
	n := data.L
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += 1.0 / data.Get(i)
	}
	return sum / float64(n)
}

// Geometric mean
func geomMean(data *Vector) float64 {
	n := data.L
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += math.Log(data.Get(i))
	}
	return math.Exp(sum / float64(n))
}

// Arithmetic mean
func arithMean(data *Vector) float64 {
	n := data.L
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += data.Get(i)
	}
	return sum / float64(n)
}

// Generalized mean
func genMean(data *Vector, p float64) float64 {
	n := data.L
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += math.Pow(data.Get(i), p)
	}
	return math.Pow(sum/float64(n), 1/p)
}
