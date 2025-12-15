package list

// Reverses the order of the contents of a slice using two pointers
func Flip(list []int, start int, stop int) {

	x, y := start, stop
	for y > x {
		list[x], list[y] = list[y], list[x]
		x++
		y--
	}

}

// Exchanges 2 sub arrays in a list based on the index passed, effectively rotating the array by that number of positions
func (list List) Rotate(sliceLength int) {

	Flip(list, 0, len(list)-1)

	Flip(list, 0, sliceLength-1)

	Flip(list, sliceLength, len(list)-1)

}
