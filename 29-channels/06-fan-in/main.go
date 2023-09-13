package main

import (
	"fmt"
	"sync"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	var xi []int

	go send(eve, odd)

	go receive(eve, odd, fanin)

	for v := range fanin {
		if len(xi) == 0 || v > xi[len(xi)-1] {
			xi = append(xi, v)
		} else {
			for j := 0; j < len(xi); j++ {
				if xi[j] > v {
					xi = append(xi[:j+1], xi[j:]...)
					xi[j] = v
					break
				}
			}
		}
	}

	fmt.Println(xi)
	fmt.Println("about to exit")
}

func receive(e, o <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fanin <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range o {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
func send(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}
