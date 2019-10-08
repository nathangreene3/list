package list

type item struct {
	value      Interface
	prev, next *item
}
