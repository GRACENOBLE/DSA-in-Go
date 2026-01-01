package main

import (
	"fmt"

	"github.com/GRACENOBLE/practice/list"
)

func main() {
	myList := list.List{1, 3, 4, 5, 6, 7}

	response := myList.BinarySearch(2)

	fmt.Println(response)
}
