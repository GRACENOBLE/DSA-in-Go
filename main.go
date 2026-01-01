package main

import (
	"fmt"

	"github.com/GRACENOBLE/practice/search"
)

func main() {
	myArray := []any{"hi", 1, true, "hello"}
	found := search.UnorderedSequentialSearch(myArray, "hello")
	fmt.Println(found)
}
