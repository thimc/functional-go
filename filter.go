package main

type filterCallback[T comparable] func(element T, index int, array []T) bool

func filter[T comparable](array []T, callback filterCallback[T]) []T {
	var out []T
	for i, element := range array {
		if callback(element, i, array) {
			out = append(out, element)
		}
	}

	return out
}

