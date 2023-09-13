package main

import "fmt"

func main() {
	fmt.Println(Sum([]int{2, 3, 1, 4}...))
}

// Sum adds an unlimited number of values of type int.
func Sum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}
