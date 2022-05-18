package main

import (
	"bytes"
	"testing"
)

func TestWordCount(t *testing.T) {
	b := bytes.NewBufferString("testing the word count function\n")
	expected := 5
	got := count(b, false)

	if expected != got {
		t.Errorf("Expected %d, got %d instead", expected, got)
	}
}

func TestLineCount(t *testing.T) {
	b := bytes.NewBufferString("Line1 Word2 Word3\nLine2\nLine3\nLine4")
	expected := 4
	got := count(b, true)

	if expected != got {
		t.Errorf("Expected %d, got %d instead", expected, got)
	}
}
