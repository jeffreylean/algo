package algo

import (
	"reflect"
	"unsafe"
)

func StringToByte(input string) []byte {
	var b []byte
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&input))
	byteHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	byteHeader.Data = strHeader.Data
	byteHeader.Cap = len(input)
	byteHeader.Len = len(input)

	return b
}
