package main

import (
	"testing"
)

func TestHamming(t *testing.T) {
	left := []byte("this is a test")
	right := []byte("wokka wokka!!!")
	distance := hamming(left, right)
	expect := 37
	if distance != expect {
		t.Errorf("expected %v\ngot %v", expect, distance)
	}
}
