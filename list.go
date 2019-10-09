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

// Find a value.
func (ls *List) Find(value Comparable) *Item {
	if ls.length == 0 {
		return nil
	}

	for itm := ls.head; itm != nil; itm = itm.next {
		if itm.Value.Compare(value) == 0 {
			return itm
		}
	}

	return nil
}

// Insert several values.
func (ls *List) Insert(values ...Comparable) {
	for _, value := range values {
		ls.insert(value)
	}
}

// insert a value.
func (ls *List) insert(value Comparable) {
	switch {
	case ls.length == 0:
		ls.head = &Item{Value: value}
		ls.tail = ls.head
	case 0 < ls.head.Value.Compare(value):
		ls.head.prev = &Item{
			Value: value,
			next:  ls.head,
		}

		ls.head = ls.head.prev
	default:
		for itm := ls.tail; itm != nil; itm = itm.prev {
			if itm.Value.Compare(value) <= 0 {
				if itm == ls.tail {
					ls.tail.next = &Item{
						Value: value,
						prev:  ls.tail,
					}

					ls.tail = ls.tail.next
				} else {
					itm.next.prev = &Item{
						Value: value,
						prev:  itm,
						next:  itm.next,
					}

					itm.next = itm.next.prev
				}

				break
			}
		}
	}

	ls.length++
}

// Length of the list.
func (ls *List) Length() int {
	return ls.length
}

// Remove several values. If duplicates exist, they will all be removed.
func (ls *List) Remove(values ...Comparable) {
	for _, value := range values {
		ls.remove(value)
	}
}

// remove a value. If duplicates exist, they will all be removed.
func (ls *List) remove(value Comparable) {
	for itm := ls.Find(value); itm != nil; itm = ls.Find(value) {
		switch {
		case ls.length == 1:
			ls.head = nil
			ls.tail = nil
		case itm == ls.head:
			ls.head = ls.head.next
		case itm == ls.tail:
			ls.tail = ls.tail.prev
		default:
			itm.prev.next = itm.next
			itm.next.prev = itm.prev
		}

		ls.length--
	}
}

// Slice comparable values.
func (ls *List) Slice() []Comparable {
	s := make([]Comparable, 0, ls.length)
	for itm := ls.head; itm != nil; itm = itm.next {
		s = append(s, itm.Value)
	}

	return s
}

// Map comparable values.
func (ls *List) Map() map[int]Comparable {
	m := make(map[int]Comparable)
	var i int
	for itm := ls.head; itm != nil; itm = itm.next {
		m[i] = itm.Value
		i++
	}

	return m
}
