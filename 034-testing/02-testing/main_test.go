package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{4, 6, 10, 22}, 42},
		{[]int{3, 7, 10, 22}, 42},
		{[]int{1, 2, 3}, 6},
		{[]int{4, 5, 7}, 16},
	}

	for _, v := range tests {
		sum := mySum(v.data...)
		if sum != v.answer {
			t.Error("Expected", v.answer, "GOT", sum)
		}
	}

}
