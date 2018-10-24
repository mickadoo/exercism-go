// Package grep finds patterns in input
package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	flagLineNum         = "-n" // Print the line numbers of each matching line
	flagCaseInsensitive = "-i" // Match line using a case-insensitive comparison.
	flagFileNameOnly    = "-l" // Print only the names of files that contain at least one matching line.
	flagFullLinesOnly   = "-x" // Only match entire lines, instead of lines that contain a match.
	flagInvert          = "-v" // Invert the program -- collect all lines that fail to match the pattern.
)

var flags []string

func isFlagEnabled(name string) bool {
	for _, flag := range flags {
		if flag == name {
			return true
		}
	}

	return false
}

// Search finds a matching patterns in a list of given files
func Search(pattern string, flagList []string, fileNames []string) []string {
	flags = flagList
	matches := make([]string, 0)
	isMultiple := len(fileNames) > 1

	for _, fileName := range fileNames {
		file := openFile(fileName)
		scanner := bufio.NewScanner(file)
		lineNum := 1

		for scanner.Scan() {
			line := scanner.Text()
			if hasMatch(line, pattern) {
				newMatch := formatMatch(line, lineNum, fileName, isMultiple)
				matches = append(matches, newMatch)

				// we can stop checking file here if we only care about the filename
				if isFlagEnabled(flagFileNameOnly) {
					break
				}
			}
			lineNum++
		}
		file.Close()
	}

	return matches
}

func hasMatch(line string, pattern string) bool {
	hasMatch := false

	if isFlagEnabled(flagCaseInsensitive) {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	if isFlagEnabled(flagFullLinesOnly) {
		hasMatch = line == pattern
	} else {
		hasMatch = strings.Contains(line, pattern)
	}

	if isFlagEnabled(flagInvert) {
		return !hasMatch
	}

	return hasMatch
}

func formatMatch(line string, lineNum int, fileName string, isMultiple bool) string {
	if isFlagEnabled(flagFileNameOnly) {
		return fileName
	}

	if isFlagEnabled(flagLineNum) {
		line = fmt.Sprintf("%d:%s", lineNum, line)
	}

	if isMultiple {
		line = fmt.Sprintf("%s:%s", fileName, line)
	}

	return line
}

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	return file
}
