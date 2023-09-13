package main

import (
	"fmt"
)

func main() {
	// // Buffered channel solution
	// c := make(chan int, 1)

	// c <- 42

	// go func solution
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
