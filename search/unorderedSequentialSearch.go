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
