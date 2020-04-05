package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type nameOff int32 // offset to a name
type typeOff int32 // offset to an *rtype
type textOff int32 // offset from top of text section

type tflag uint8

type rtype struct {
	size       uintptr
	ptrdata    uintptr // number of bytes in the type that can contain pointers
	hash       uint32  // hash of type; avoids computation in hash tables
	tflag      tflag   // extra type information flags
	align      uint8   // alignment of variable with this type
	fieldAlign uint8   // alignment of struct field with this type
	kind       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal     func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata    *byte   // garbage collection data
	str       nameOff // string form
	ptrToThis typeOff // type for pointer to this type, may be zero
}

type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}

func main() {
	a := 1000                                    // variable is int, but variable doesnt have any information about type, calculates on compilation time
	f := interface{}(a)                          // boxing variable to empty interface, inteface is boxed(container) type
	ei := *(*emptyInterface)(unsafe.Pointer(&f)) // show interface box type
	fmt.Println(ei.typ)
	fmt.Println(reflect.TypeOf(&a)) // we cannot retriev type without boxing variable
}
