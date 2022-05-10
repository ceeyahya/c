package main

import (
	"bytes"
	"testing"
)

func TestWordCount(t *testing.T) {
	b := bytes.NewBufferString("testing the word count function\n")
	expected := 5
	got := count(b)

	if expected != got {
		t.Errorf("Expected %d, got %d instead", expected, got)
	}
}
