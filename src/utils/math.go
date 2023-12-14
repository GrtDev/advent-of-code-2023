package utils

func PowInt(x int, y int) int {
	if y == 0 {
		return 1
	}
	result := x
	for i := 2; i <= y; i++ {
		result *= x
	}
	return result
}

func AbsInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func CombinationsInt(n int, r int) int {

	result := 1

	for i := 0; i < r; i++ {
		result *= (n - i)
		result /= (i + 1)
	}

	return result
}
