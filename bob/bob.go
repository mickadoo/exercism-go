// Package bob includes functionality for bob, a chatbot
package bob

import (
	"strings"
	"unicode"
)

// Hey responds to text input with a comment
func Hey(remark string) string {
	var isQuestion, isUpper bool

	remark = strings.TrimSpace(remark)

	isQuestion = strings.HasSuffix(remark, "?")
	isAlpha := strings.IndexFunc(remark, unicode.IsLetter) >= 0
	isUpper = remark == strings.ToUpper(remark) && isAlpha

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isQuestion && !isUpper:
		return "Sure."
	case isQuestion:
		return "Calm down, I know what I'm doing!"
	case isUpper:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
