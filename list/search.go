package list

// -----------------------Sequential search---------------------

func (list List) SequentialSearch( searchTerm int) Response {
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


// -----------------------Binary search---------------------

func (list List) BinarySearch(searchTerm int) Response {
	if len(list) == 0 {
		return Response{Presence: false, Number: searchTerm}
	}

	var currentSample []int = list
	for len(currentSample) > 0 {

		middle := GetMiddle(currentSample)

		if middle.Value == searchTerm {
			return Response{Presence: true, Number: searchTerm}
		} else if middle.Value > searchTerm {
			currentSample = currentSample[:middle.Index]
		} else {
			currentSample = currentSample[middle.Index+1:]
		}
	}
	return Response{Presence: false, Number: searchTerm}
}

type MiddleValue struct {
	Index int
	Value int
}

func GetMiddle(sublist []int) MiddleValue {
	middleFloorIndex := int(len(sublist) / 2)
	return MiddleValue{
		Index: middleFloorIndex,
		Value: sublist[middleFloorIndex],
	}
}

