package main

import (
	"fmt"
)

func main() {
	// this doesn't work because you can't receive from a send only channel
	// cs := make(chan<- int)
	// this doesn't work because you can't send to a receive only channel
	// cs := make(<-chan int)
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)

	fmt.Printf("--------\n")
	fmt.Printf("cs\t%T\n", cs)
}
