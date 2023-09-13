package main

import "fmt"

func main() {

	xi := []int{1, 41, 23, 41}
	fmt.Println(MySum(xi...))
}

// Adds an unlimited number of values of type int
func MySum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}
