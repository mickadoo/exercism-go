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

type flagSettings []string

var flags flagSettings

func (f flagSettings) isEnabled(name string) bool {
	for _, flag := range f {
		if flag == name {
			return true
		}
	}

	return false
}

// Search finds a matching patterns in a list of given files
func Search(pattern string, flagList []string, fileNames []string) []string {
	var matches = make([]string, 0)
	var flags = flagSettings(flagList)
	var isMultiple = len(fileNames) > 1

	for _, fileName := range fileNames {
		file := openFile(fileName)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 1
		for scanner.Scan() {
			line := scanner.Text()
			if hasMatch(line, pattern, flags) {
				matches = append(matches, formatMatch(line, lineNum, fileName, flags, isMultiple))

				// we can stop checking file here if we only care about the filename
				if flags.isEnabled(flagFileNameOnly) {
					break
				}
			}
			lineNum++
		}
	}

	return matches
}

func hasMatch(line string, pattern string, flags flagSettings) bool {
	var hasMatch = false

	if flags.isEnabled(flagCaseInsensitive) {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	if flags.isEnabled(flagFullLinesOnly) {
		hasMatch = line == pattern
	} else {
		hasMatch = strings.Contains(line, pattern)
	}

	if flags.isEnabled(flagInvert) {
		return !hasMatch
	}

	return hasMatch
}

func formatMatch(line string, lineNum int, fileName string, flags flagSettings, isMultiple bool) string {
	if flags.isEnabled(flagLineNum) {
		line = fmt.Sprintf("%d:%s", lineNum, line)
	}

	if isMultiple {
		line = fmt.Sprintf("%s:%s", fileName, line)
	}

	if flags.isEnabled(flagFileNameOnly) {
		line = fileName
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
