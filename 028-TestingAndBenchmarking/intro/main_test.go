package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected ", 5, "Got", x)
	}
}

func TestSquare(t *testing.T) {
	x := square(3)
	if x != 9 {
		t.Error("Expected", 9, "Got", x)
	}
}
