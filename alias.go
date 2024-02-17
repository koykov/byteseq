package byteseq

// Q is a shorthand alias of Byteseq.
type Q = Byteseq

// Q2S is a shorthand alias of SequenceToString().
func Q2S[T Q](x T) string {
	return SequenceToString(x)
}

// Q2B is a shorthand alias of SequenceToBytes().
func Q2B[T Q](x T) []byte {
	return SequenceToBytes(x)
}

// S2Q is a shorthand alias of StringToSequence().
func S2Q[T Q](s string) T {
	return StringToSequence[T](s)
}

// B2Q is a shorthand alias of BytesToSequence().
func B2Q[T Q](p []byte) T {
	return BytesToSequence[T](p)
}
