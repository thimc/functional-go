package main

type mapfCallback[T any] func(element T) T

func mapf[T any](array []T, callback mapfCallback[T]) []T {
	var out []T
	for _, element := range array {
		out = append(out, callback(element))
	}

	return out
}
