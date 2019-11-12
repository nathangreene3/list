package sortedlist

// item ...
type item struct {
	Value      Comparable
	prev, next *item
}
