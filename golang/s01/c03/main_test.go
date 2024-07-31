package main

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	var input string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	result, err := xor(s1, s2)
	if err != nil {
		t.Errorf("%v", err)
	}
	if expect != result {
		t.Errorf("expected %v\ngot %v", expect, result)
	}
}
