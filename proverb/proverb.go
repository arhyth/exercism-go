// Package proverb exposes Proverb function
package proverb

import "fmt"

// Proverb receives a slice of strings and returns a proverb string
func Proverb(rhyme []string) []string {
	proverb := []string{}
	length := len(rhyme)

	if length > 1 {
		for i := 0; i < (length - 1); i++ {
			sentence :=  fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i + 1])
			proverb = append(proverb, sentence)
		}
	}
	
	if length > 0 {
		sentence := fmt.Sprintf("And all for the want of a %v.", rhyme[0])
		proverb = append(proverb, sentence)
	}
	
	return proverb
}
