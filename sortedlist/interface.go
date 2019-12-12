package sortedlist

// Comparable defines how a value should be compared within the context of a sorted list.
type Comparable interface {
	Compare(c Comparable) int
}
