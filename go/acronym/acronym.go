// Package acronym creates acronyms.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate creates acronyms based on strings in the input.
func Abbreviate(s string) string {
	reg := regexp.MustCompile("[-\\s]")
	split := reg.Split(s, -1)
	coll := make([]string, 10)

	for idx, chunk := range split {
		coll[idx] = strings.ToUpper(string(chunk[0]))
	}

	return strings.Join(coll, "")
}