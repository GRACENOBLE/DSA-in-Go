// UNORDERED SEQUENTIAL SEARCH
// Time complexity O(n) as the time increases with size of the array for worst case
// Space complexity O(1) as no more space is needed other than defining the array

package search

type response struct {
	found bool
	value any
}

func UnorderedSequentialSearch(array []any, searchTerm any) response {
	for _, value := range array {
		if value == searchTerm {
			return response{
				found: true,
				value: value,
			}
		}
	}

	return response{
		found: false,
		value: nil,
	}
}
