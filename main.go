package main

import (
	"fmt"

	"github.com/GRACENOBLE/practice/sum"
)

func main() {

	myList := []int{-5, -4, -3, -2, -1}
	fmt.Println(sum.LargestContiguousSum(myList))
}
