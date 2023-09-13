package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(eve, odd, quit)

	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func receive(e, o <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the eve channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok:", i, ok)
				return
			} else {
				fmt.Println("from comma ok :", i)
			}
		}
	}
}
func send(e, o chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(quit)
}
