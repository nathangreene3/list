package sortedlist

// item ...
type item struct {
	Value      Interface
	prev, next *item
}
