// Package hamming exposes the Distance function.
package hamming

import "errors"

// Distance receives two string arguments and returns an int, error
// tuple. The int value indicating the number of different characters
// between the two strings.
func Distance(a, b string) (int, error) {
	runes_a := []rune(a)
	runes_b := []rune(b)
	count := 0
	if len(runes_a) != len(runes_b) {
		return 0, errors.New("Strings are not of equal length")
	}
	for i, c := range runes_a {
		if runes_b[i] != c {
			count += 1	
		}
	}
	return count, nil
}
