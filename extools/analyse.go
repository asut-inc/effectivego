package extools

import (
	"fmt"
	"reflect"
)

func AnalyseData(v interface{}) {
	// expect CustomStruct if non pointer
	objType := reflect.TypeOf(v)
	g := reflect.ValueOf(objType).Interface()
	fmt.Println(g)
	//fmt.Printf("show type: %v \n", g)
	//switch g {
	//case "*[]int":
	//	fmt.Println("test")
	////case "t":
	////	fmt.Printf("slice: len= ")
	//	//header := (*reflect.SliceHeader)(unsafe.Pointer(&tet))
	//	//spew.Dump(header)
	//default:
	//	fmt.Println("Unknown type", g)
	//}
}
