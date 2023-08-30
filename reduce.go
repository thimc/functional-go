package main

type reduceCallback[T comparable] func(accumulator T, current T) T

func reduce[T comparable](array []T, callback reduceCallback[T], initial T) T {
	var out T = initial
	for _, element := range array {
		out = callback(element, out)
	}

	return out
}
