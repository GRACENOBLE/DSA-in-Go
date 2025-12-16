package list

func (list List) SumIntegers() int {
	var sum int

	for _, number := range list {
		sum += number
	}

	return sum
}
