// Package twofer is used for generating string to explain sharing of items
package twofer

import "fmt"

// ShareWith takes a name and returns a string in the form "One for $name, one for me"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
