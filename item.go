package list

// item holds a value and references it's previous and next items, if any.
type item struct {
	value      interface{}
	prev, next *item
}
