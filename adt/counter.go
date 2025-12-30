// Counter used to keep track of the count of items.
package adt

type Counter map[any]int

// Add an element
func (c *Counter) Add(key any) {
	(*c)[key] += 1
}

// Find an element
func (c Counter) Find(key any) bool {
	_, ok := c[key]
	return ok
}

// Get the count of an elemet by its Key
func (c Counter) Get(key any) int {
	value := c[key]
	return value
}
