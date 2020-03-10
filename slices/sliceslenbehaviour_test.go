package slices

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func copySliceHeader(sl []int) {
	sl = append(sl, 999)
	fmt.Println("Copied slice", sl)
}

func TestSlicesLenBehaviour(t *testing.T) {
	sl := make([]int, 1, 10)
	fmt.Println("Initialized slice", sl)

	copySliceHeader(sl)
	fmt.Println("Inited slice", sl)

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&sl))
	sliceHeader.Len = 2
	sl = *(*[]int)(unsafe.Pointer(sliceHeader))
	fmt.Println("Inited slice with changed len", sl)
}
