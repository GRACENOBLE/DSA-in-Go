// This is used when we do not know whether the data in the provided list is sorted at all.

package search

// import "slices"

type Response struct {
	Presence bool
	Number   int
}

func SearchNumberInList(list []int, searchTerm int) Response {
	for _, number := range list {
		if number == searchTerm {
			return Response{
				Presence: true,
				Number:   searchTerm,
			}
		}
	}
	
	// There is an inbuilt function in go to check for the presence of an item
	// if slices.Contains(list, searchTerm) {
	// 	return Response{
	// 		Presence: true,
	// 		Number:   searchTerm,
	// 	}
	// }

	return Response{
		Presence: false,
		Number:   searchTerm,
	}
}
