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

func xor(s1 string, s2 string) (string, error) {
	b1, err := hex.DecodeString(s1)
	if err != nil {
		return "", err
	}
	b2, err := hex.DecodeString(s2)
	if err != nil {
		return "", err
	}
	if len(b1) != len(b2) {
		return "", fmt.Errorf("lengths do not match")
	}
	out := make([]byte, len(b1))
	for i, _ := range b1 {
		out[i] = b1[i] ^ b2[i]
	}
	return hex.EncodeToString(out), nil
}
