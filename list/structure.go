package list

type List []int

// Response struct for what is returned when you find something in a list
type Response struct {
	Presence bool
	Number   int
}
