package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 10)
	}
}

func TestSubtract(t *testing.T) {
	total := Subtract(4, 2)
	if total != 2 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 2)
	}
}

func TestDoMath(t *testing.T) {
	total := DoMath(2, 3, Add)
	if total != 5 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 5)
	}

	total = DoMath(3, 2, Subtract)
	if total != 1 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 1)
	}
}
