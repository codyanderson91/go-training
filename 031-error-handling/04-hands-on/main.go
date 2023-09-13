package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishees", "Never say never"},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("An Error occuered :", err)
	}

	fmt.Println(string(bs))
}
