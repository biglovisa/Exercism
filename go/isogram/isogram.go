package isogram

import (
	"strings"
	"regexp"
	"log"
)


func IsIsogram(testString string) bool {
	isIso := true
	found := make(map[string]bool)

	// Make regexp for only keeping letters
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
			log.Fatal(err)
	}

	testString = strings.ToLower(reg.ReplaceAllString(testString, ""))

	for _, letter := range strings.Split(testString, "") {
		if found[letter] {
			isIso = false
			return isIso
		}
		found[letter] = true
	}

	return isIso
}