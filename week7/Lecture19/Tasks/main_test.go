package main

import (
	"io"
	"testing"
)

func TestRead(t *testing.T) {
	toBeReversed := NewReverseStringReader("blabla")

	buf := make([]byte, 6)

	for {
		_, err := toBeReversed.Read(buf)

		if err == io.EOF {
			break
		}
	}

	if string(buf) != "albalb" {
		t.Errorf("Expected \"albalb\", got %s", string(buf))
	}

}
