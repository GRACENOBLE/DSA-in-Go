package adt

import (
	"fmt"

	"github.com/golang-collections/collections/splay"
	"github.com/golang-collections/collections/trie"
	"github.com/golang-collections/collections/tst"
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

// Stores index and value pairs
func Trie() {
	t := trie.New()

	t.Insert(1, 5)
	t.Insert(2, 4)
	t.Insert(3, 3)
	t.Insert(4, 2)
	t.Insert(5, 1)

	// Check if value exists
	found := t.Has(5)
	fmt.Println("Found 5:", found)

	// Get an item at a specific index
	item := t.Get(4)
	fmt.Println("Found item at 4:", item)

	// Iterate over the items in the structure
	t.Do(func(k, v interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
		return true
	})
}

// Ternary search tree (stores key(string) and value pairs)
func Tst() {
	t := tst.New()

	t.Insert("one", 5)
	t.Insert("two", 4)
	t.Insert("three", 3)
	t.Insert("four", 2)
	t.Insert("five", 1)

	// Check for presence
	present := t.Has("one")
	fmt.Println("Has one: ", present)

	// Get a value for a specific key
	item := t.Get("one")
	fmt.Println("The item at key one is:", item)

	// Remove an item from the tree
	t.Remove("five")

	// Iterate through the values of the tree
	t.Do(func(k string, v interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
		return true
	})
}
