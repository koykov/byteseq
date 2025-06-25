package byteseq

import (
	"unsafe"
)

// Byteseq represents generic byte sequence types.
type Byteseq interface {
	~string | ~[]byte
}

// SequenceToString makes fast conversion of sequence to string.
func SequenceToString[T Byteseq](x T) string {
	ix := any(x)
	switch y := ix.(type) {
	case string:
		return y
	case []byte:
		return *(*string)(unsafe.Pointer(&y))
	default:
		return ""
	}
}

// SequenceToBytes makes fast conversion of sequence to bytes.
func SequenceToBytes[T Byteseq](x T) []byte {
	ix := any(x)
	switch y := ix.(type) {
	case string:
		sh := (*sheader)(unsafe.Pointer(&y))
		var h header
		h.Data = sh.Data
		h.Len = sh.Len
		h.Cap = sh.Len
		return *(*[]byte)(unsafe.Pointer(&h))
	case []byte:
		return y
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

// ToBytes tries to convert x to bytes.
func ToBytes[T Byteseq](x T) (p []byte, ok bool) {
	p, ok = any(x).([]byte)
	return
}

// ToString tries to convert x to string.
func ToString[T Byteseq](x T) (s string, ok bool) {
	s, ok = any(x).(string)
	return
}
