package list

// Interface ...
type Interface interface {
	Comparable
}

// Comparable ...
type Comparable interface {
	Compare(c Comparable) int
}
