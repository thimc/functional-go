package main

import (
	"reflect"
	"testing"
)

func BenchmarkMapf(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func TestMapfNumbers(t *testing.T) {
	numbers := []int{1, 4, 9, 16}
	expected := []int{2, 8, 18, 32}
	got := mapf[int](numbers, func(element int) int {
		return element * 2
	})

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %+v, got %+v", expected, got)
	}
}

func TestMapfSideEffect(t *testing.T) {
	cart := []float64{5,15,25}
	total := 0.0

	withTax := mapf[float64](cart, func(element float64) float64 {
		total += element
		return element * 1.2
	})
	expectedTax := []float64{6.0,18.0,30.0}
	expectedTotal := 45.0

	if total != expectedTotal {
		t.Fatalf("expected total %f, got %f", expectedTotal, total)
	}

	if !reflect.DeepEqual(withTax, expectedTax) {
		t.Fatalf("expected tax %+v, got %+v", expectedTax, withTax)
	}
}
