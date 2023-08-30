package main

import (
	"reflect"
	"testing"
)

func BenchmarkFilter(b * testing.B) {
	for i := 0; i < b.N; i++ {
		filter[int]([]int{1,2,3,4,5,6,7,8,9,10}, func(element, index int, array []int) bool {
			return element > 1
		})
	}
}

func BenchmarkFilterUsingFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var out []int
		for _, element := range []int{1,2,3,4,5,6,7,8,9,10} {
			if element > 1 {
				out = append(out, element)
			}
		}
	}
}

func TestFilterNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := []int{2, 3, 4, 5}

	filtered := filter[int](numbers, func(element, _ int, _ []int) bool {
		return element > 1
	})

	if !reflect.DeepEqual(filtered, expected) {
		t.Fatalf("expected the result to be %+v, got %+v", expected, filtered)
	}
}

func TestFilterStructs(t *testing.T) {
	type User struct {
		Name       string
		Age        int
		Occupation string
	}
	users := []User{
		{Name: "John", Age: 25, Occupation: "gardener"},
		{Name: "Lenny", Age: 51, Occupation: "programmer"},
		{Name: "Andrew", Age: 43, Occupation: "teacher"},
		{Name: "Peter", Age: 81, Occupation: "teacher"},
		{Name: "Anna", Age: 47, Occupation: "programmer"},
		{Name: "Albert", Age: 76, Occupation: "programmer"},
	}
	expected := []User {
		users[1], users[4], users[5],
	}

	filtered := filter[User](users, func(user User, _ int, _ []User) bool {
		return user.Age > 40 && user.Occupation == "programmer"
	})

	if !reflect.DeepEqual(filtered, expected) {
		t.Fatalf("expected the struct to be %+v, got %+v", expected, filtered)
	}
}
