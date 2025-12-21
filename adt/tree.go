package adt

import (
	"fmt"

	"github.com/golang-collections/collections/splay"
)

func less(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func Splay() {
	s := splay.New(less)

	s.Add(5)
	s.Add(3)
	s.Add(7)
	s.Add(1)
	s.Add(9)

	// Check if a value exists
	found := s.Has(7)
	fmt.Println("Found 7:", found)

	// Get the first item (min)
	first := s.First()
	fmt.Println("First:", first)

	// Get the last (max)
	last := s.Last()
	fmt.Println("Last:", last)

	// Delete an element in the tree
	s.Remove(3)

	// Iterate through the elements in the tree
	s.Do(func(v interface{}) bool {
		fmt.Println(v)
		return true
	})

}

func Trae() {

}
