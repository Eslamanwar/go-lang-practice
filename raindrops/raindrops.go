package raindrops

import (
	"fmt"
)

func rainDrops(i int) string {
	var result string

	if i%3 == 0 {
		result += "Pling"
	}
	if i%5 == 0 {
		result += "Plang"
	}
	if i%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		return fmt.Sprint(i)
	}
	return result

}
