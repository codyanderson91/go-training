package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return []byte{}, fmt.Errorf("Error: %v", err)
	}

	return bs, err
}
