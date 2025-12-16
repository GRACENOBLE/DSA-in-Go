// Given a list of positive and negative integers, find a contiguous subarray whose sum (sum of elements) is maximum.

package list

// Kabanes algorithm, maintains the highest value of maxSoFar
func (list List) LargestContiguousSum() int {

	maxSoFar := list[0]
	maxEndingHere := list[0]

	for _, number := range list[1:] {
		maxEndingHere = max(maxEndingHere+number, number)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}
