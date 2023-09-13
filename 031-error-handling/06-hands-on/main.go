package main

import (
	"errors"
	"fmt"
)

type location struct {
	err error
}

func (l location) Error() string {
	return fmt.Sprintf("This is my customer error: %v", l.err)
}

func main() {

	_, err := notNegative(-10.23)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("This code is over")
}

func notNegative(f float64) (float64, error) {
	if f < 0 {
		e := errors.New("Why did you pass in a negative to this function")
		return 0, location{err: e}
	}

	return 42, nil
}
