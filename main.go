package main

import (
	"fmt"

	"github.com/GRACENOBLE/practice/adt"
)

func main() {
	myHashMap := make(adt.Hashmap)
	myHashMap.Insert("one", 1)
	myHashMap.Insert("two", 2)
	myHashMap.Insert("three", 3)
	myHashMap.Insert("four", 4)

	myHashMap.Delete("four")

	fmt.Println(myHashMap)
	fmt.Println(myHashMap.Search("three"))
}
