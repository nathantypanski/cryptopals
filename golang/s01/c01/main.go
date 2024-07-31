package main

import (
	"fmt"
	"encoding/base64"
	"encoding/hex"
)

func hexToBase64(s string) (string, error) {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(decoded), nil
}

func main() {
	var hex string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	if out, err := hexToBase64(hex); err != nil {
		fmt.Println(out)
	} else {
		fmt.Errorf("got error: %v", err)
	}
}
