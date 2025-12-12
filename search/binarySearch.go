// The assumption is that the list is sorted.

package search

func (list List) BinarySearch(searchTerm int) Response {
	if len(list.Items) == 0 {
		return Response{Presence: false, Number: searchTerm}
	}

	var currentSample []int = list.Items
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
