package byteseq

type sheader struct {
	Data uintptr
	Len  int
}

type header struct {
	Data uintptr
	Len  int
	Cap  int
}
