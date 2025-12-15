package main

import (
	"fmt"

	"github.com/GRACENOBLE/practice/list"
)

func main() {

	integers := list.List{1, 2, 3, 4, 5, 6, 7}

	integers.Rotate(3)

	fmt.Printf("%v", integers)
}
