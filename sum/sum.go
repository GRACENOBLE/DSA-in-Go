package sum

func SumIntegers(arr []int) int {
	var sum int

	for _, number := range arr {
		sum += number
	}

	return sum
}
