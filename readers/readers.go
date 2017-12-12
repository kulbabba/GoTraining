package readers

import (
	"bytes"
	"io"
)

// Combine returns an io.Reader which represents
// the contents of a and b.
func Combine(a, b io.Reader) io.Reader {

	buf := new(bytes.Buffer)
	buf.ReadFrom(a)
	buf.ReadFrom(b)
	return buf

}

// always reader always fills the read buffer with
// the byte ch.
type alwaysReader struct {
	ch byte
}

func (a *alwaysReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = a.ch
	}
	return len(buf), nil
}

// AReader returns an io.Reader which returns n 'A' characters
func AReader(n int) io.Reader {

	buf := new(bytes.Buffer)
	for i := 1; i <= n; i++ {
		buf.WriteString("A")
	}
	return buf
}
