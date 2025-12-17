// A recursive puzzle that uses theree stacks(rods) to move n disks to a destination rod using a temporary rod as storage, this has the condition that a bigger disk is never placed on top of a smaller disk

package recursiveproblems

import "fmt"

func TOHUtil(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}
	TOHUtil(num-1, from, temp, to)
	fmt.Println("Move disk", num, "from peg", from, "to peg", to)
	TOHUtil(num-1, temp, to, from)
}

func TowersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in Tower of Hanoi are:")
	TOHUtil(num, "A", "C", "B")
}
