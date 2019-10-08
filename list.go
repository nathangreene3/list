package list

// List ...
type List struct {
	head, tail *item
	length     int
}

// New ...
func New(values ...Comparable) *List {
	var ls List
	ls.Insert(values...)
	return &ls
}

// Insert ...
func (ls *List) Insert(values ...Comparable) {
	for _, value := range values {
		ls.insert(value)
	}
}

// insert ...
func (ls *List) insert(value Comparable) {
	newItm := item{value: value}
	if ls.length == 0 {
		// New list
		ls.head = &newItm
		ls.tail = &newItm
	} else {
		var inserted bool
		for itm := ls.head; itm != nil && !inserted; itm = itm.next {
			if 0 < itm.value.Compare(value) {
				if itm == ls.head {
					ls.head = &newItm
				} else {
					itm.prev.next = &newItm
				}

				itm.prev = &newItm
				newItm.next = itm
				inserted = true
			}
		}

		if !inserted {
			newItm.prev = ls.tail
			ls.tail.next = &newItm
			ls.tail = &newItm
		}
	}

	ls.length++
}

// Slice ...
func (ls *List) Slice() []Comparable {
	s := make([]Comparable, 0, ls.length)
	for itm := ls.head; itm != nil; itm = itm.next {
		s = append(s, itm.value)
	}

	return s
}
