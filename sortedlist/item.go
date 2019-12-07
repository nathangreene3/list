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

// compare an item's value to a given value.
func (itm *item) compare(value Comparable) int {
	return itm.Value.Compare(value)
}

// contains returns true if an item contains a value.
func (itm *item) contains(value Comparable) bool {
	return itm.Value.Compare(value) == 0
}

// copy an item.
func (itm *item) copy() *item {
	return &item{Value: itm.Value, prev: itm.prev, next: itm.next}
}

// equals returns true if two items contain the same value.
func (itm *item) equals(item *item) bool {
	return itm.Value.Compare(item.Value) == 0
}
