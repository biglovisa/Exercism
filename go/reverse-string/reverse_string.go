package reverse

import (
	"strings"
)

func Reverse(word string) string {
	reversedWord := ""

	for _, str := range strings.Split(word, "") {
		reversedWord = string(str) + reversedWord
	}

	return string(reversedWord)
}