package main

import (
	"bytes"
	"io"
)

// ReadAll reads all the lines of text from r and returns
// all the data read as a string
func ReadAll(r io.Reader) string {
	// TODO: add solution here
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	s := buf.String()
	return s
}
