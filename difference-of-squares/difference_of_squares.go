package diffsquares

import (
	"math"
)

//SquareOfSum fn
func SquareOfSum(number int) int {
	var sum int
	for i := 1; i <= number; i++ {
		sum += i
	}
	return int((math.Pow(float64(sum), 2)))

}

//SumOfSquares fn
func SumOfSquares(number int) int {
	var sumOfSquares float64
	for i := 1; i <= number; i++ {
		sumOfSquares += math.Pow(float64(i), 2)
	}

	return int(sumOfSquares)
}

//Difference fn
func Difference(number int) int {

	return (SquareOfSum(number) - SumOfSquares(number))
}
