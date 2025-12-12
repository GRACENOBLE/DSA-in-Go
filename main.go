package main

import (
	"fmt"

	"github.com/GRACENOBLE/practice/search"
)

func main() {
	integers := []int{1, 2, 3, 4, 5}

	response := search.SearchNumberInList(integers, 5)

	fmt.Printf("%v", response)
}
