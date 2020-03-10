package unsafe

import (
	"reflect"
	"unsafe"
)

func strToBytes(s string) []byte {
	header := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytesHeader := &reflect.SliceHeader{
		Data: header.Data,
		Len:  header.Len,
		Cap:  header.Len,
	}
	return *(*[]byte)(unsafe.Pointer(bytesHeader))
}
