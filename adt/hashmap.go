// Maps dont store elements in contiguous order

package adt

type Hashmap map[string]int

func (m Hashmap) Insert(key string, value int) {
	m[key] = value
}

func (m Hashmap) Delete(key string) {
	delete(m, key)
}

type FoundObject struct {
	IsPresent bool
	value     int
}

func (m Hashmap) Search(key string) FoundObject {
	value, ok := m[key]
	return FoundObject{IsPresent: ok, value: value}
}
