// The stack ADT from go collections 
// Methods Push() Pop() Top() Size() IsEmpty()
package adt

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func Stack() {
	s := stack.New()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	for s.Len() != 0 {
		val := s.Pop()
		fmt.Println(val)
	}
}
