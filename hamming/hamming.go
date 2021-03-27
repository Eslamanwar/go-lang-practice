package hamming

import (
	"errors"
	"fmt"
)

// Distance between hamming
func Distance(a, b string) (int, error) {

	var hammingDistance int
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hammingDistance++
			}
		}
		return hammingDistance, nil

	}
	return 0, errors.New("Must be the same size")

}

func main() {
	i, err := Distance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT")
	fmt.Println(i, err)
}
