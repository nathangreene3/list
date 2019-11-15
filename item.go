package list

// item ...
type item struct {
	Value      interface{}
	prev, next *item
}

func newItem(value interface{}, prev, next *item) *item {
	return &item{Value: value, prev: prev, next: next}
}

// getFrom returns the ith item from a starting item.
func (itm *item) getFrom(i int) *item {
	ithItm := itm
	for ; 0 < i; ithItm = ithItm.next {
		i--
	}

	return ithItm
}
