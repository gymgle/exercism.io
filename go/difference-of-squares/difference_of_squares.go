package diffsquares

// Difference the difference between the square of the sum
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SquareOfSum with a standard algorithm
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares with a standard algorithm
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}
