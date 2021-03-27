package isogram

import (
	"strings"
)

// IsIsogram fn
func IsIsogram(word string) bool {
	for i, v := range word {
		if string(v) == "-" || string(v) == " " {
			continue
		}

		index := i + 1
		if strings.Contains(strings.ToUpper(word[index:]), strings.ToUpper(string(v))) {

			return false
		}
	}
	return true
}
