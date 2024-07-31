package main

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	var s1 string = "1c0111001f010100061a024b53535009181c"
	var s2 string = "686974207468652062756c6c277320657965"
	var expect string = "746865206b696420646f6e277420706c6179"
	result, err := xor(s1, s2)
	if err != nil {
		t.Errorf("%v", err)
	}
	if expect != result {
		t.Errorf("expected %v\ngot %v", expect, result)
	}
}
