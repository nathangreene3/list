package list

import (
	"container/heap"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var (
	seeded bool
	seed   int64
)

func initSeed() {
	if !seeded {
		seed = time.Now().Unix()
		rand.Seed(seed)
		seeded = true
	}
}

// TestList ensures manipulating a list is equivalent to manipulating a slice.
func TestList(t *testing.T) {
	initSeed()

	var (
		iters    = 100
		numItems = 100
	)

	for i := 0; i < iters; i++ {
		ls := New(func(x, y interface{}) bool { return x.(int) < y.(int) })
		nums := make([]int, 0, numItems)
		for j := 0; j < numItems; j++ {
			x := rand.Int()
			ls.Append(x)
			nums = append(nums, x)
		}

		var (
			itm = ls.head
			j   int
		)

		for ; j < numItems && itm != nil && itm.next != nil; itm = itm.next {
			index, ok := ls.Search(nums[j])
			if j != index || !ok {
				fmt.Printf("seed: %d\n", seed)
				t.Fatalf("\nexpected (%d, %t)\nreceived (%d, %t)\n", j, true, index, ok)
			}

			j++
		}

		if s := ls.Slice(); len(s) != ls.length {
			fmt.Printf("seed: %d\n", seed)
			t.Fatalf("\nexpected length %d\nreceived %d\n", ls.length, len(s))
		}
	}
}

func TestInsertRemove(t *testing.T) {
	var (
		s  = []int{0, 1, 2, 3, 4, 5, 6}
		ls = New(func(x, y interface{}) bool { return x.(int) < y.(int) }, 0, 1, 0, 2, 0)
	)

	ls.RemoveAt(4)
	ls.RemoveAt(2)
	ls.RemoveAt(0)
	ls.Append(0, 4, 0, 5, 0).Remove(0).InsertAt(2, 3).InsertAt(0, 0).Append(6)
	if len(s) != ls.Len() {
		t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
	}

	for i := 0; i < len(s); i++ {
		index, ok := ls.Search(s[i])
		if i != index || !ok {
			t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
		}
	}
}

func TestSort(t *testing.T) {
	var (
		s  = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		ls = New(func(x, y interface{}) bool { return x.(int) < y.(int) }, 9, 0, 8, 1, 7, 2, 6, 3, 5, 4).Sort()
	)

	if len(s) != ls.Len() {
		t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
	}

	for i := 0; i < len(s); i++ {
		index, ok := ls.Search(s[i])
		if i != index || !ok {
			t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
		}
	}
}

func TestHeapSort(t *testing.T) {
	var (
		s0 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		ls = New(func(x, y interface{}) bool { return x.(int) < y.(int) })
	)

	heap.Push(ls, 9)
	heap.Push(ls, 0)
	heap.Push(ls, 8)
	heap.Push(ls, 1)
	heap.Push(ls, 7)
	heap.Push(ls, 2)
	heap.Push(ls, 6)
	heap.Push(ls, 3)
	heap.Push(ls, 5)
	heap.Push(ls, 4)

	s1 := []int{
		heap.Pop(ls).(int),
		heap.Pop(ls).(int),
		heap.Pop(ls).(int),
		heap.Pop(ls).(int),
		heap.Pop(ls).(int),
		heap.Pop(ls).(int),
		heap.Pop(ls).(int),
		heap.Pop(ls).(int),
		heap.Pop(ls).(int),
		heap.Pop(ls).(int),
	}

	if len(s0) != len(s1) {
		t.Fatalf("\nexpected %v\nreceived %v\n", s0, s1)
	}

	for i := 0; i < len(s0); i++ {
		if s0[i] != s1[i] {
			t.Fatalf("\nexpected %v\nreceived %v\n", s0, s1)
		}
	}
}
