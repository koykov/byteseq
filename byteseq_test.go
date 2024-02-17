package byteseq

import (
	"bytes"
	"strings"
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
	t.Run("to bytes", func(t *testing.T) {
		in := bytes.Repeat([]byte("foobar "), 100)
		out, ok := ToBytes(in)
		if !ok && !bytes.Equal(in, out) {
			t.FailNow()
		}
	})
	t.Run("to string", func(t *testing.T) {
		in := strings.Repeat("foobar ", 100)
		out, ok := ToString(in)
		if !ok && in != out {
			t.FailNow()
		}
	})
}

func BenchmarkByteseq(b *testing.B) {
	b.Run("sequence2string", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := Q2S("foobar")
			_ = r
		}
	})
	b.Run("sequence2bytes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := Q2B([]byte("foobar"))
			_ = r
		}
	})
	b.Run("string2sequence", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := S2Q[string]("foobar")
			_ = r
		}
	})
	b.Run("bytes2sequence", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := B2Q[[]byte]([]byte("foobar"))
			_ = r
		}
	})
	b.Run("to bytes", func(b *testing.B) {
		b.ReportAllocs()
		in := bytes.Repeat([]byte("foobar "), 100)
		for i := 0; i < b.N; i++ {
			r, ok := ToBytes(in)
			_, _ = r, ok
		}
	})
	b.Run("to string", func(b *testing.B) {
		b.ReportAllocs()
		in := strings.Repeat("foobar ", 100)
		for i := 0; i < b.N; i++ {
			r, ok := ToString(in)
			_, _ = r, ok
		}
	})
}
