package main

import (
	"testing"
)

func BenchmarkReduceSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce[int](
			benchmarkArray,
			func(accumulator, current int) int { return accumulator + current },
			0,
		)
	}
}

func BenchmarkReduceSumFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, element := range benchmarkArray {
			sum += element
		}
	}
}

func TestReduceWithoutInitialValue(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	initial := 0
	expected := 10

	sum := reduce[int](
		numbers,
		func(accumulator, current int) int { return accumulator + current },
		initial,
	)

	if sum != expected {
		t.Fatalf("expected the sum to be %d, got %d", expected, sum)
	}
}

func TestReduceWithInitialValue(t *testing.T) {
	numbers := []int{15, 16, 17, 18, 19}
	initial := 10
	expected := 95

	sum := reduce[int](
		numbers,
		func(accumulator, current int) int { return accumulator + current },
		initial,
	)

	if sum != expected {
		t.Fatalf("expected the sum to be %d, got %d", expected, sum)
	}
}
