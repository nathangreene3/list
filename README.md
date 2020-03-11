# List

```go
go get github.com/nathangreene3/list
```

A `List` is a doubly-linked list of `interface{}` values.

## Sorted List

```go
type Comparable interface {
    Compare(c Comparable) int
}
```

A `SortedList` is a sorted, doubly-linked list of `Comparable` values.
