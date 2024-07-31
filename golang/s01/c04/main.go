package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func hexToBase64(s string) (string, error) {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(decoded), nil
}

func xor(b1 []byte, b2 []byte) (string, error) {
	if len(b1) != len(b2) {
		return "", fmt.Errorf("lengths do not match")
	}
	out := make([]byte, len(b1))
	for i, _ := range b1 {
		out[i] = b1[i] ^ b2[i]
	}
	return hex.EncodeToString(out), nil
}

// https://crypto.stackexchange.com/questions/30209/developing-algorithm-for-detecting-plain-text-via-frequency-analysis
//
// chi squared -> https://en.wikipedia.org/wiki/Chi-squared_test
func score(input string) float64 {
	freq := []float64{
		0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, // A-G
		0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, // H-N
		0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, // O-U
		0.00978, 0.02360, 0.00150, 0.01974, 0.00074, // V-Z
	}

	var count = make([]int, 26)
	ignored := 0
	for i := range count {
		count[i] = 0
	}
	for _, e := range input {
		c := rune(e)
		if c >= 65 && c <= 90 {
			count[c-65]++ // uppercase A-Z
		} else if c >= 97 && c <= 122 {
			count[c-97]++ // lowercase a-z
		} else if c >= 32 && c <= 126 {
			ignored++ // numbers and punct.
		} else if c == 9 || c == 10 || c == 13 {
			ignored++ // TAB, CR, LF
		} else {
			return -1 // not printable ASCII = impossible(?)
		}

	}

	var chi2 float64 = 0.0
	len := len(input) - ignored
	for i, _ := range freq {
		observed := float64(count[i])
		expected := float64(len) * freq[i]
		difference := observed - expected
		chi2 += difference * difference / expected
	}
	return chi2
}

func xorSingleChar(s []byte, c rune) (string, error) {
	repeated := make([]byte, len(s))
	for i, _ := range repeated {
		repeated[i] = byte(c)
	}
	return xor(s, repeated)
}

func findBestXor(input []byte) (rune, string, error) {
	first := int(rune('0'))
	last := int(rune('Z'))
	scoreLen := last - first
	scores := make([]float64, scoreLen)
	for i := 0; i < scoreLen; i++ {
		out, _ := xorSingleChar(input, rune(first + i))
		decoded, err := hex.DecodeString(out)
		if err != nil {
			return rune(0), "", fmt.Errorf("%v while decoding '%v' with '%v'", err, out)
		}
		out = string(decoded)
		scores[i] = score(out)
	}

	min := 0
	c := rune(first)
	decoded, err := xorSingleChar(input, c)
	if err != nil {
		return rune(0), "", fmt.Errorf("%v while xoring '%v' with '%v'", err, c, input)
	}
	for i := 0; i < scoreLen; i++ {
		if (scores[i] < scores[min]) && (scores[i] != -1) || (scores[min] == -1) {
			min = i
			c = rune(first + i)
			out, err := xorSingleChar(input, c)
			if err != nil {
				return rune(0), "", fmt.Errorf("%v while xoring '%v' with '%v'", err, c, input)
			}
			bytes, err := hex.DecodeString(out)
			decoded = string(bytes)
			if err != nil {
				return rune(0), "", fmt.Errorf("%v while decoding '%v' with '%v'", err, out)
			}
		}
	}

	return c, decoded, nil
}

func findBestString(inputs [][]byte) (rune, string, string, error) {
	if len(inputs) == 0 {
		return rune(0), "", "", fmt.Errorf("no input strings\n")
	}

	c, decoded, err := findBestXor(inputs[0])
	if err != nil {
		return rune(0), "", "", fmt.Errorf("Error finding best XOR: %v\n", err)
	}

	bestLine := inputs[0]
	bestScore := score(decoded)
	for _, s := range inputs {
		if proposedC, proposedDecoded, err := findBestXor(s); err == nil {
			proposedLine := s
			proposedScore := score(proposedDecoded)
			if proposedScore < bestScore && proposedScore != -1 || bestScore == -1 {
				c = proposedC
				decoded = proposedDecoded
				bestLine = proposedLine
				bestScore = proposedScore
			}
		} else if err != nil {
			return rune(0), "", "", fmt.Errorf("Error scoring best XOR: %v\n", err)
		}
	}
	return c, decoded, string(bestLine), nil
}

func main() {
	var lines [][]byte = [][]byte{}
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(f)

	// Iterate over each line in the file
	for scanner.Scan() {
		decoded, err := hex.DecodeString(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error decoding string: %v\n", err)
			os.Exit(1)
		}
		lines = append(lines, decoded)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	c, decoded, bestLine, err := findBestString(lines)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error finding best: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v (repeated) ^ %v = %v\n", string(c), bestLine, decoded)
}
