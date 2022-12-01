package main

import (
	"io"
	"os"
)

type ReverseStringReader struct {
	input string
	i     int64
}

func (rsr *ReverseStringReader) Read(b []byte) (n int, err error) {
	if rsr.i >= int64(len(rsr.input)) {
		return 0, io.EOF
	}

	n = copy(b, reverseSlice(rsr.input[rsr.i:]))
	rsr.i += int64(n)
	return
}

func reverseSlice(s string) []byte {
	data := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func NewReverseStringReader(input string) *ReverseStringReader {

	return &ReverseStringReader{input: input}
}

func main() {
	toBeReversed := NewReverseStringReader("My name is Deyan.")
	io.Copy(os.Stdout, toBeReversed)
}
