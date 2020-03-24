package raindrops

import (
	"strconv"
)

func Convert(num int) string {
	value := ""
	if num % 3 == 0 {
		value = "Pling"
	}
	if num % 5 == 0 {
		value = value + "Plang"
	}
	if num % 7 == 0{
		value = value + "Plong"
	}
	if value == "" {
		value = strconv.Itoa(num)
	}

	return value
}