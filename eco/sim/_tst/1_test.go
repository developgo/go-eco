package sim

import (
	"testing"
	//	"fmt"
	"code.google.com/p/go-eco/eco/aux"
)

// test against known values
func TestEuclid2(t *testing.T) {
	var (
		data, out *Matrix
		d, s      float64
	)

	data = NewMatrix(2, 3)
	data.Set(0, 0, 0)
	data.Set(0, 1, 0)
	data.Set(0, 2, 0)
	data.Set(1, 0, 1)
	data.Set(1, 1, 1)
	data.Set(1, 2, 1)

	out = Euclid_D(data)
	d = 1.7320508075688771
	s = 0.36602540378443865

	if !check(out.Get(0, 0), out.Get(1, 1)) {
		t.Error()
	}
	if !check(out.Get(0, 0), 0.0) {
		t.Error()
	}
	if !check(out.Get(0, 1), out.Get(1, 0)) {
		t.Error()
	}
	if !check(out.Get(0, 1), d) {
		t.Error()
	}

	out = Euclid_S(data)

	if !check(out.Get(0, 0), out.Get(1, 1)) {
		t.Error()
	}
	if !check(out.Get(0, 0), 1.0) {
		t.Error()
	}
	if !check(out.Get(0, 1), out.Get(1, 0)) {
		t.Error()
	}
	if !check(out.Get(0, 1), s) {
		t.Error()
	}

	out = Manhattan_D(data)
	d = 3
	s = 0.25

	if !check(out.Get(0, 0), out.Get(1, 1)) {
		t.Error()
	}
	if !check(out.Get(0, 0), 0.0) {
		t.Error()
	}
	if !check(out.Get(0, 1), out.Get(1, 0)) {
		t.Error()
	}
	if !check(out.Get(0, 1), d) {
		t.Error()
	}

	out = Manhattan_S(data)

	if !check(out.Get(0, 0), out.Get(1, 1)) {
		t.Error()
	}
	if !check(out.Get(0, 0), 1.0) {
		t.Error()
	}
	if !check(out.Get(0, 1), out.Get(1, 0)) {
		t.Error()
	}
	if !check(out.Get(0, 1), s) {
		t.Error()
	}

}

// test distances against known values
func TestD(t *testing.T) {
	var (
		data, out *Matrix
		d         float64
	)

	/*data:
	4.637511 5.795001 5.700484 6.524882 6.170708 6.690082 6.156994 6.921186 4.740336 5.400573
	7.535951 5.804745 6.697524 5.297671 5.77213 7.187614 7.470511 6.559553 5.870524 5.31025
	4.428564 5.698517 4.882601 6.541425 4.910434 7.311253 4.562559 6.858137 7.105823 5.963177
	6.614128 4.282033 6.146613 5.819527 6.797518 5.657332 5.61468 5.180996 5.374655 6.594351
	6.79283 6.371214 5.990534 6.518546 6.283301 6.841622 5.978732 5.278547 7.825815 6.36177
	4.813066 6.990308 6.809527 7.83582 6.256215 4.981545 7.230944 5.322504 5.981109 5.738691
	*/

	rows := 6
	cols := 10
	arr := [...]float64{4.637511, 5.795001, 5.700484, 6.524882, 6.170708, 6.690082, 6.156994, 6.921186, 4.740336, 5.400573, 7.535951, 5.804745, 6.697524, 5.297671, 5.77213, 7.187614, 7.470511, 6.559553, 5.870524, 5.31025, 4.428564, 5.698517, 4.882601, 6.541425, 4.910434, 7.311253, 4.562559, 6.858137, 7.105823, 5.963177, 6.614128, 4.282033, 6.146613, 5.819527, 6.797518, 5.657332, 5.61468, 5.180996, 5.374655, 6.594351, 6.79283, 6.371214, 5.990534, 6.518546, 6.283301, 6.841622, 5.978732, 5.278547, 7.825815, 6.36177, 4.813066, 6.990308, 6.809527, 7.83582, 6.256215, 4.981545, 7.230944, 5.322504, 5.981109, 5.738691}

	data = NewMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data.Set(i, j, arr[i*cols+j])

		}
	}

	// Gower
	out = Gower_D(data)
	d = 0.4322016
	if !check(out.Get(0, 5), d) {
		t.Error()
	}
	if !check(out.Get(5, 0), d) {
		t.Error()
	}
	d = 0.625486
	if !check(out.Get(3, 2), d) {
		t.Error()
	}
	if !check(out.Get(2, 3), d) {
		t.Error()
	}

	d = 0.4161464
	if !check(out.Get(4, 5), d) {
		t.Error()
	}
	if !check(out.Get(5, 4), d) {
		t.Error()
	}
	for i := 0; i < rows; i++ {
		if !check(out.Get(i, i), 0.0) {
			t.Error()
		}
	}

}
