package typestest

import (
	"reflect"
	"testing"
)

/**
* slice is struct type with unsafe pointer
* https://golang.org/src/runtime/slice.go
*	type slice struct {
*		array unsafe.Pointer
*		len   int
*		cap   int
*	}
 */

func constructSlice() []int {
	return []int{1, 2, 3, 4}
}

// slices is not pointer type
// when pass as value arguments its copies header struct
func sliceByValueAppend(a []int) []int {
	// in passed struct happens new array allocation and pointer in this header changes to new pointer
	a = append(a, 5)

	// return new struct with new pointer
	return a
}

func TestSliceByValueAppend(t *testing.T) {
	a := constructSlice()
	b := sliceByValueAppend(a) // returned new struct from function, with new pointer
	// comparison
	if reflect.DeepEqual(a, b) {
		t.Errorf("Slices equal A: %v, B: %v", a, b)
	}
}

// when passing slice as pointer, its passes pointer to its struct
func sliceByPointerAppend(a *[]int) {
	b := *a
	b = append(b, 5)

	a = &b

	return
}

func TestSliceByPointerAppend(t *testing.T) {
	a := constructSlice()
	b := a
	sliceByPointerAppend(&a)

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Slices equal A: %v, B: %v", a, b)
	}
}
