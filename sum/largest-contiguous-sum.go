// Given a list of positive and negative integers, find a contiguous subarray whose sum (sum of elements) is maximum.

package sum

// Kabanes algorithm, maintains the highest value of maxSoFar
func LargestContiguousSum(list []int) int {

	maxSoFar := list[0]
	maxEndingHere := list[0]

	for _, number := range list[1:] {
		maxEndingHere = max(maxEndingHere+number, number)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}
