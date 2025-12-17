// --------------- Euclid's Algorithm ---------------

// Finding the greatest common divisor given a pair of numbers

package recursiveproblems

func GCD(m int, n int) int {
	if m < n {
		return GCD(n, m)
	}

	if m%n == 0 {
		return n
	}

	return GCD(n, m%n)
}
