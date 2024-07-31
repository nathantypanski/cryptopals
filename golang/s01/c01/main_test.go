package main

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	var hex string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var expect string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, err := hexToBase64(hex)
	if err != nil {
		t.Errorf("%v", err)
	}
	if expect != result {
		t.Errorf("expected %v\ngot %v", expect, result)
	}
}
