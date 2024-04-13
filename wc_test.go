package main

import (
	"bytes"
	"testing"
)

func TestCountWord(t *testing.T) {
	t.Run("Counts a set of words", func(t *testing.T) {
		expected := 4

		b := bytes.NewBufferString("word word word word\n")
		got := count(b)
		if expected != got {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}
