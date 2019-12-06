package sortedlist

// item holds a comparable value and references the previous and next items.
type item struct {
	Value      Comparable
	prev, next *item
}

// newItem returns a new item.
func newItem(value Comparable, prev, next *item) *item {
	return &item{Value: value, prev: prev, next: next}
}
