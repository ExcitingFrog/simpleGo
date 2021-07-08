package convert

import (
	"reflect"
	"unsafe"
)

// unsafe
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// unsafe
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
