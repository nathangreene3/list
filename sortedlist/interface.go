package sortedlist

// Comparable ...
type Comparable interface {
	Compare(c Comparable) int
}
