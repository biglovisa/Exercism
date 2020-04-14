package luhn

import (
	"strconv"
)

func Valid(seq string) bool {
	// reverse string, do not include whitespaces
	r_str := ""
	for i := 0; i < len(seq); i++ {
		if string(seq[i]) == " " {
			continue
		}
		r_str = string(seq[i]) + r_str
	}

	// ensure length is greater than 1 character
	if len(r_str) < 2 {
		return false
	}

	// perform the actual luhn algorithm
	sum := 0
	for i := 0; i < len(r_str); i++ {
		digit, err := strconv.Atoi(string(r_str[i]));
		if err != nil {
			return false
		}

		if i % 2 != 0 {
			digit = digit * 2
			if digit >= 9 {
				digit = digit - 9
			}
		}

		sum += digit
	}

	// verify that the sum of the sum is evenly divisble by 10
	return sum % 10 == 0
}
