package main

import (
	"fmt"
	"sync"
)

type person struct {
	first string
	last  string
}
type animal struct {
	first string
	last  string
}

type saysThings interface {
	saySomething()
}

func (p *person) saySomething() {
	fmt.Println("Hi my name is", p.first)
}

func (a animal) saySomething() {
	fmt.Println("Hi I am a", a.first)
}

func talk(s saysThings) {
	s.saySomething()
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	p2 := person{
		first: "Harry",
		last:  "Barry",
	}
	p3 := person{
		first: "King",
		last:  "Kong",
	}

	a1 := animal{
		first: "Cat",
		last:  "Gray",
	}
	a2 := animal{
		first: "Dog",
		last:  "Black",
	}

	wg.Add(5)

	go talk(&p1)
	go talk(&p2)
	go talk(&p3)
	go talk(a1)
	go talk(a2)

	wg.Wait()
}
