// Min heap each node is smaller than or equal to its children, in a max heap it is the opposite

package adt

// slice to store the integers
type IntHeap []int

// Returns the length of the heap
func (h IntHeap) Len() int {
	return len(h)
}

// Returns the less of two values by index in the heap
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Switches the position of two elemens in the heap
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Adds a new elwment into a heap
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Remove the root node in the heap
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
