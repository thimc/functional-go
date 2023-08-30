package main

import "testing"

func BenchmarkFind(b * testing.B) {
	for i := 0; i < b.N; i++ {
		find[int](benchmarkArray, func(element int) bool {
			return element > 1
		})
	}
}
func BenchmarkFindFor(b * testing.B) {
	for i := 0; i < b.N; i++ {
		for _, element := range benchmarkArray {
			if element > 1 {
				break
			}
		}
	}
}

func TestFindNumber(t *testing.T) {
	numbers := []int{5, 12, 8, 130, 44}
	expected := 12

	found := find(numbers, func(element int) bool {
		return element > 10
	})

	if found != expected {
		t.Fatalf("expected found value %d, got %d", expected, found)
	}
}

func TestFindStruct(t *testing.T) {
	type Item struct {
		Name     string
		Quantity int
	}
	inventory := []Item{
		{Name: "apples", Quantity: 2},
		{Name: "bananas", Quantity: 0},
		{Name: "cherries", Quantity: 5},
	}
	expected := Item{
		Name: "cherries", Quantity: 5,
	}

	isCherry := func(element Item) bool {
		return element.Name == expected.Name
	}

	found := find[Item](inventory, isCherry)

	if found.Name != expected.Name{
		t.Fatalf("expected name %s, got %s", expected.Name, found.Name)
	}
	if found.Quantity != expected.Quantity{
		t.Fatalf("expected quantity %d, got %d", expected.Quantity, found.Quantity)
	}
}
