package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	// c <- 42
	// fmt.Println(<-c)
	// c <- 43
	go func() {
		c <- 42
		c <- 43
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
