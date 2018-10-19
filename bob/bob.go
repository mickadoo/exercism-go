// Package bob includes functionality for bob, a chatbot
package bob

import (
	"strings"
)

func endsWith(str string, char rune) bool {
	length := len(str)
	if length == 0 {
		return false
	}
	return rune(str[length-1]) == char
}

// Hey responds to text input with a comment
func Hey(remark string) string {
	var isQuestion, isUpper bool

	// trim whitespace characters
	remark = strings.Trim(remark, " \t\r\n")

	isQuestion = endsWith(remark, '?')
	isAlpha := strings.ContainsAny(strings.ToLower(remark), "abcdefghijklmnopqrstuvwxyz")
	isUpper = remark == strings.ToUpper(remark) && isAlpha

	switch true {
	case remark == "":
		return "Fine. Be that way!"
	case isQuestion && !isUpper:
		return "Sure."
	case isQuestion && isUpper:
		return "Calm down, I know what I'm doing!"
	case isUpper && !isQuestion:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
