package main

import (
	"fmt"
)

func main() {
	numbers := []int{505, 268, 110, 13, 50, 16, 49, 19, 100, 64}
	fmt.Printf("%+v -> filter -> map -> reduce = ", numbers)

	fmt.Printf("%d\n",
		reduce[int](filter[int](mapf[int](numbers,
			func(element int) int { return element * 3 }),
			func(element, _ int, _ []int) bool { return element%2 == 0 }),
			func(accumulator, current int) int { return accumulator + current }, -487))
}
