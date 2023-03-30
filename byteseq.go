package byteseq

import (
	"reflect"
	"unsafe"
)

// Byteseq represents generic byte sequence types.
type Byteseq interface {
	~string | ~[]byte
}

// SequenceToString makes fast conversion of sequence to string.
func SequenceToString[T Byteseq](x T) string {
	ix := any(x)
	switch ix.(type) {
	case string:
		return ix.(string)
	case []byte:
		p := ix.([]byte)
		return *(*string)(unsafe.Pointer(&p))
	default:
		return ""
	}
}

// SequenceToBytes makes fast conversion of sequence to bytes.
func SequenceToBytes[T Byteseq](x T) []byte {
	ix := any(x)
	switch ix.(type) {
	case string:
		s := ix.(string)
		sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
		var h reflect.SliceHeader
		h.Data = sh.Data
		h.Len = sh.Len
		h.Cap = sh.Len
		return *(*[]byte)(unsafe.Pointer(&h))
	case []byte:
		return ix.([]byte)
	default:
		return nil
	}
}

// StringToSequence makes fast conversion of string to sequence.
func StringToSequence[T Byteseq](s string) T {
	return *(*T)(unsafe.Pointer(&s))
}

// BytesToSequence makes fast conversion of bytes to sequence.
func BytesToSequence[T Byteseq](p []byte) T {
	return *(*T)(unsafe.Pointer(&p))
}
