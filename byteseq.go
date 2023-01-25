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

// Q2S is a shorthand alias of SequenceToString().
func Q2S[T Byteseq](x T) string {
	return SequenceToString(x)
}

// Q2B is a shorthand alias of SequenceToBytes().
func Q2B[T Byteseq](x T) []byte {
	return SequenceToBytes(x)
}
