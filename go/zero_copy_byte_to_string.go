package algo

import (
	"unsafe"
)

func StringToByte(input string) []byte {
	strPtr := unsafe.StringData(input)
	b := unsafe.Slice(strPtr, len(input))

	return b
}
