package encode

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func writeEncodedOutput(buffer *bytes.Buffer, char rune, count int) {
	if count == 1 {
		buffer.WriteRune(char)
	} else {
		buffer.WriteString(fmt.Sprintf("%d%c", count, char))
	}
}

// RunLengthEncode encodes the input using run-length encoding
func RunLengthEncode(input string) string {
	inputBuffer := strings.NewReader(input)
	outputBuffer := bytes.Buffer{}
	currentCount := 1

	for {
		current, _, cErr := inputBuffer.ReadRune()
		next, _, nErr := inputBuffer.ReadRune()

		if cErr != nil {
			break
		}
		if nErr != nil {
			writeEncodedOutput(&outputBuffer, current, currentCount)
			break
		}
		inputBuffer.UnreadRune()
		if next != current {
			writeEncodedOutput(&outputBuffer, current, currentCount)
			currentCount = 1
		} else {
			currentCount++
		}
	}

	return outputBuffer.String()
}

// RunLengthDecode decodes the run-length-encoded input
func RunLengthDecode(input string) string {
	inputBuffer := strings.NewReader(input)
	outputBuffer := bytes.Buffer{}
	var currentCountStr string

	for {
		current, _, cErr := inputBuffer.ReadRune()
		if cErr != nil {
			break
		}
		if unicode.IsNumber(current) {
			currentCountStr += string(current)
			continue
		} else {
			if currentCountStr == "" {
				outputBuffer.WriteRune(current)
			} else {
				currentCount, err := strconv.Atoi(currentCountStr)
				if err == nil {
					outputBuffer.WriteString(strings.Repeat(string(current), currentCount))
					currentCountStr = ""
				}
			}
		}
	}

	return outputBuffer.String()
}
