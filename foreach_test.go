package main

import "testing"

func TestForEach(t *testing.T) {
	runes := []rune{'a','b','c'}
	ok := 0

	forEach[rune](runes, func(element rune) {
		// t.Log(string(element))
		ok++
	})

	if ok != len(runes) {
		t.Fatalf("expected %d iterations, had %d", len(runes), ok)
	}
}
