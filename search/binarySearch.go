// SORTED SEQUENTIAL SEARCH
// If the list is sorted then we do not have to traverse all the items in the list, this will save us a lot of time and make the algorithm more efficient
// We always discard half the list if we dont find the item in the current iteration
// Time complexity O(log n)
package search

func (l List) BinarySearch(searchTerm int) Response {
	size := len(l)
	low := 0
	high := size - 1
	mid := 0

	for low < high {
		mid = low + (high-low)/2
		if l[mid] == searchTerm {
			return Response{
				found: true,
				value: searchTerm,
			}
		} else if l[mid] < searchTerm {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return Response{
		found: false,
		value: searchTerm,
	}
}

func findMiddle(sl []int) int {
	return sl[len(sl)/2]
}
