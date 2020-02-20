package extools

import (
	"testing"
)

func TestAnalyseData(t *testing.T) {
	slice := []int{32, 23, 23, 2, 3, 23}
	//header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	ch := make(chan int, 3)
	AnalyseData(&ch)
	AnalyseData(&slice)
}
