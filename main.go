package main

import (
	"fmt"

	"github.com/GRACENOBLE/practice/search"
)

func main() {

	integers := search.List{Items: []int{1, 3, 5, 7}}

	response := integers.BinarySearch(6)

	// response := search.GetMiddle(integers.Items)

	fmt.Printf("%v", response)
}

