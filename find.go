package main

type findCallback[T any] func(element T) bool

func find[T any](array []T, callback findCallback[T]) T {
	var out T
	for _, element := range array {
		if callback(element) {
			return element
		}
	}

	return out
}
