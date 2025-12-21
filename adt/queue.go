// The queue ADT from go collections
// Methods Enqueue() Dequeue() Peek() Size() IsEmpty()

package adt

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

func Queue() {
	q := queue.New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	for q.Len() != 0 {
		val := q.Dequeue()
		fmt.Println(val)
	}
}
