//Package reverse is involved in reversing things
package reverse

// String returns the reverse of a string
func String(str string) string {
	var res string

	for _, char := range str {
		res = string(char) + res
	}

	return string(res)
}
