package main

type forEachCallback[T any] func(element T)

func forEach[T any](array []T, callback forEachCallback[T]) {
	for _, element := range array {
		callback(element)
	}
}
