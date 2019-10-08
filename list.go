package list

// List ...
type List struct {
	head, tail *Item
	length     int
}

// New ...
func New(values ...Comparable) *List {
	var ls List
	ls.Insert(values...)
	return &ls
}

// Find ...
func (ls *List) Find(value Comparable) *Item {
	for itm := ls.head; itm != nil; itm = itm.next {
		if itm.value.Compare(value) == 0 {
			return itm
		}
	}

	return nil
}

// Insert ...
func (ls *List) Insert(values ...Comparable) {
	for _, value := range values {
		ls.insert(value)
	}
}

// insert ...
func (ls *List) insert(value Comparable) {
	newItm := Item{value: value}
	if ls.length == 0 {
		ls.head = &newItm
		ls.tail = &newItm
		ls.length = 1
	} else {
		var inserted bool
		for itm := ls.head; itm != nil && !inserted; itm = itm.next {
			if 0 < itm.value.Compare(value) {
				if itm.prev != nil {
					itm.prev.next = &newItm
				} else {
					ls.head = &newItm
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

		ls.length++
	}
}

func (ls *List) remove(value Comparable) {

}

// Slice ...
func (ls *List) Slice() []Comparable {
	s := make([]Comparable, 0, ls.length)
	for itm := ls.head; itm != nil; itm = itm.next {
		s = append(s, itm.value)
	}

	return s
}
