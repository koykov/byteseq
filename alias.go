package byteseq

// Q2S is a shorthand alias of SequenceToString().
func Q2S[T Byteseq](x T) string {
	return SequenceToString(x)
}

// Q2B is a shorthand alias of SequenceToBytes().
func Q2B[T Byteseq](x T) []byte {
	return SequenceToBytes(x)
}

// S2Q is a shorthand alias of StringToSequence().
func S2Q[T Byteseq](s string) T {
	return StringToSequence[T](s)
}

// B2Q is a shorthand alias of BytesToSequence().
func B2Q[T Byteseq](p []byte) T {
	return BytesToSequence[T](p)
}
