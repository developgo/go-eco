/* 
Circle product similarity
	   Marquardt, W.H. 1978 Archaeological seriation. In: Schiffer, M.B.(ed.)
	   Advances in Archaeological Method and Theory. Academic Press, N.Y., p.281.
	   Kendall, D.G. 1971b Seriation from abundance matrices. In: Hodson, F.R.,
	   Kendall, D.G. & Tautu, P: Mathematics in the archaeological and historical
	   sciences. Edinburgh University Press, pp. 215-252. 
*/
package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

func circleProduct_S(data *DenseMatrix) *DenseMatrix {
	var (
		i, j, k int
		sum     float64
		sim     *DenseMatrix
	)

	sim = Zeros(data.Rows(), data.Rows()) // square similarity matrix row vs. row
	i = 0
	j = 0
	k = 0
	sum = 0.0

	for i = 0; i < data.Rows(); i++ {
		sim.Set(i, i, 1.0)
	}

	for i = 0; i < data.Rows(); i++ {
		for j = i + 1; j < data.Rows(); j++ {
			sum = 0
			for k = 0; k < data.Cols(); k++ {
				sum += Min(data.Get(i, k), data.Get(j, k))
			}
			sim.Set(i, j, sum)
			sim.Set(j, i, sum)
		}
	}
	return sim
}
