package myculc

import "testing"

func TestPlus(t *testing.T) {
	sum := plus(10, 20)
	if sum != 30 {
		t.Error("Error")
	}
}

func TestSubstract(t *testing.T) {
	s := substract(20, 10)
	if s != 10 {
		t.Error("Error")
	}
}
