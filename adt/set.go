// Sores unique values, not contiguous

package adt

import (
	"fmt"

	"github.com/golang-collections/collections/set"
)

func Set() {
	// Create a new set
	st := set.New()

	// Add values to the set
	st.Insert(1)
	st.Insert(2)
	st.Insert(3)
	st.Insert(3)

	fmt.Println(st)

	// Check for presence
	fmt.Println(st.Has(3))

	// Iterate over set values
	st.Do(func(i interface{}) {
		fmt.Println(i)
	})
}
