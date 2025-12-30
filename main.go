package main

import (
	"fmt"

	"github.com/GRACENOBLE/practice/adt"
)

func main() {
	myCounter := make(adt.Counter)
	myCounter.Add("one")
	myCounter.Add("two")
	myCounter.Add("one")

	fmt.Println(myCounter.Find("two"))
	fmt.Println(myCounter.Get("one"))

}
