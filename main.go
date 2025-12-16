package main

import (
	"fmt"

	"github.com/GRACENOBLE/practice/list"
)

func main() {
	myList := list.List{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(myList.LargestContiguousSum())
}
