package typestest

import (
	"reflect"
	"testing"
)

// slices is not pointer type
// when pass as value arguments its copies header struct
func sliceByValueAppend(a []int) []int {
	a = append(a, 5)

	return a
}

func constructSlice() []int {
	return []int{1, 2, 3, 4}
}

func TestSliceByValueAppend(t *testing.T) {
	a := constructSlice()
	b := sliceByValueAppend(a)
	if reflect.DeepEqual(a, b) {
		t.Errorf("Slices equal A: %v, B: %v", a, b)
	}
}
