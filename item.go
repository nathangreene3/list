package list

// item ...
type item struct {
	Value      interface{}
	prev, next *item
}
