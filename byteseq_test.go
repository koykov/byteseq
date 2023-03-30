package byteseq

import (
	"bytes"
	"testing"
)

func TestByteseq(t *testing.T) {
	var (
		s = "foobar"
		b = []byte("foobar")
	)
	t.Run("sequence2string", func(t *testing.T) {
		r := Q2S("foobar")
		if r != s {
			t.FailNow()
		}
	})
	t.Run("sequence2bytes", func(t *testing.T) {
		r := Q2B([]byte("foobar"))
		if !bytes.Equal(r, b) {
			t.FailNow()
		}
	})
	t.Run("string2sequence", func(t *testing.T) {
		r := S2Q[string]("foobar")
		if r != s {
			t.FailNow()
		}
	})
	t.Run("bytes2sequence", func(t *testing.T) {
		r := B2Q[[]byte]([]byte("foobar"))
		if !bytes.Equal(r, b) {
			t.FailNow()
		}
	})
}

func BenchmarkByteseq(b *testing.B) {
	var (
		s = "foobar"
		p = []byte("foobar")
	)
	b.Run("sequence2string", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := Q2S("foobar")
			if r != s {
				b.FailNow()
			}
		}
	})
	b.Run("sequence2bytes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := Q2B([]byte("foobar"))
			if !bytes.Equal(r, p) {
				b.FailNow()
			}
		}
	})
	b.Run("string2sequence", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := S2Q[string]("foobar")
			if r != s {
				b.FailNow()
			}
		}
	})
	b.Run("bytes2sequence", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := B2Q[[]byte]([]byte("foobar"))
			if !bytes.Equal(r, p) {
				b.FailNow()
			}
		}
	})
}
