// Package proverb generates proverbs
package proverb

import "fmt"

// Proverb takes a list of words and outputs a proverb for it
func Proverb(rhyme []string) []string {
	length := len(rhyme)
	proverb := make([]string, len(rhyme))

	// Loop only as long as there is another word left in the list
	for index := 0; index+1 < length; index++ {
		proverb[index] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[index], rhyme[index+1])
	}

	if length > 0 {
		proverb[length-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	}

	return proverb
}
