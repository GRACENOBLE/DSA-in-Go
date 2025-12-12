package main

import (
	"fmt"

	"github.com/GRACENOBLE/practice/sum"
)

func main() {
	integers := []int{1, 2, 3, 4, 5}

	sum := sum.SumIntegers(integers)

	fmt.Printf("The sum of the integers is %v", sum)

}

