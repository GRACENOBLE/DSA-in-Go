package main

import (
	"container/heap"
	"fmt"

	"github.com/GRACENOBLE/practice/adt"
)

func main() {
	h := &adt.IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Pop(h)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
