// Package bob is about greeting people.
package bob

import (
	"log"
	"regexp"
	"strings"
)

// Hey is the main function of this program.
func Hey(remark string) string {
	cleanRemark := strings.TrimSpace(remark)

	if cleanRemark == "" {
		return "Fine. Be that way!"
	}

	if strings.ToUpper(cleanRemark) == remark && !onlyNumbers(cleanRemark) {
		return "Whoa, chill out!"
	}

	if strings.LastIndex(cleanRemark, "?") == (len(cleanRemark) - 1) {
		return "Sure."
	}

	return "Whatever."
}

func onlyNumbers(remark string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	formattedString := reg.ReplaceAllString(remark, "")
	c := 0
	for _, char := range formattedString {
		if char >= '0' && char <= '9' {
			c++
		}
	}
	return c == len(formattedString)
}