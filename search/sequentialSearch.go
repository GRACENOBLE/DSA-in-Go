// UNORDERED SEQUENTIAL SEARCH
// Time complexity O(n) as the time increases with size of the array for worst case
// Space complexity O(1) as no more space is needed other than defining the array

package search

func (l List) UnorderedSequentialSearch(searchTerm int) Response {
	for _, value := range l {
		if value == searchTerm {
			return Response{
				found: true,
				value: value,
			}
		}
	}

	return Response{
		found: false,
		value: nil,
	}
}
