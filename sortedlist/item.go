package sortedlist

// item ...
type item struct {
	value      Comparable
	prev, next *item
}
