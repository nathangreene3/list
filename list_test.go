package list

import (
	"container/heap"
	golist "container/list"
	"fmt"
	"math/rand"
	"sort"
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
		var (
			ls   = New(func(x, y interface{}) bool { return x.(int) < y.(int) })
			nums = make([]int, 0, numItems)
		)

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
			if index, ok := ls.Search(nums[j]); j != index || !ok {
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
	ls.Append(0, 4, 0, 5, 0).Remove(0).InsertAt(2, 3).Prepend(0).Append(6)
	if len(s) != ls.Len() {
		t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
	}

	for i := 0; i < len(s); i++ {
		if index, ok := ls.Search(s[i]); i != index || !ok {
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
		if index, ok := ls.Search(s[i]); i != index || !ok {
			t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
		}
	}
}

func TestHeap(t *testing.T) {
	var (
		rnd = []int{9, 0, 8, 1, 7, 2, 6, 3, 5, 4}
		exp = make([]int, len(rnd))
		ls  = New(func(x, y interface{}) bool { return x.(int) < y.(int) })
	)

	copy(exp, rnd)
	sort.Ints(exp)

	for i := 0; i < len(rnd); i++ {
		heap.Push(ls, rnd[i])
	}

	rec := make([]int, 0, len(exp))
	for 0 < ls.Len() {
		rec = append(rec, heap.Pop(ls).(int))
	}

	if len(exp) != len(rec) {
		t.Fatalf("\nexpected %v\nreceived %v\n", exp, rec)
	}

	for i := 0; i < len(exp); i++ {
		if exp[i] != rec[i] {
			t.Fatalf("\nexpected %v\nreceived %v\n", exp, rec)
		}
	}
}

func TestReduce(t *testing.T) {
	var (
		n   = 5
		exp = n * (n + 1) / 2 // 1+2+...+n
		rec = Generate(
			n,
			func(i int) interface{} { return i + 1 },
			func(x, y interface{}) bool { return x.(int) < y.(int) },
		).Reduce(func(x, y interface{}) interface{} { xVal, _ := x.(int); yVal, _ := y.(int); return xVal + yVal })
	)

	if exp != rec {
		t.Fatalf("\nexpected %d\nreceived %d\n", exp, rec)
	}
}

func BenchmarkList(b *testing.B) {
	var (
		maxSize = int(256)
		values  = make([]interface{}, maxSize)
		less    = func(x, y interface{}) bool { return x.(int) < y.(int) }
	)

	{ // Linear benchmark
		stepSize := 8
		for i := 0; i <= len(values); i += stepSize {
			if benchmarkList(b, less, values[:i]...) {
				// b.Error("An unexpected error occured")
			}
		}

		for i := 0; i <= len(values); i += stepSize {
			if benchmarkGoList(b, values[:i]...) {
				// b.Error("An unexpected error occured")
			}
		}

		for i := 0; i <= len(values); i += stepSize {
			if benchmarkSliceCopy(b, values[:i]...) {
				// b.Error("An unexpected error occured")
			}
		}

		for i := 0; i <= len(values); i += stepSize {
			if benchmarkSliceAppend(b, values[:i]...) {
				// b.Error("An unexpected error occured")
			}
		}
	}

	{ // Exponential benchmark
		for i := 1; i <= maxSize; i <<= 1 {
			if benchmarkList(b, less, values[:i]...) {
				// b.Error("An unexpected error occured")
			}
		}

		for i := 1; i <= maxSize; i <<= 1 {
			if benchmarkGoList(b, values[:i]...) {
				// b.Error("An unexpected error occured")
			}
		}

		for i := 1; i <= maxSize; i <<= 1 {
			if benchmarkSliceCopy(b, values[:i]...) {
				// b.Error("An unexpected error occured")
			}
		}

		for i := 1; i <= maxSize; i <<= 1 {
			if benchmarkSliceAppend(b, values[:i]...) {
				// b.Error("An unexpected error occured")
			}
		}
	}
}

func benchmarkList(b *testing.B, less Less, values ...interface{}) bool {
	r := b.Run(
		fmt.Sprintf("New list of %d values", len(values)),
		func(b0 *testing.B) {
			for i := 0; i < b0.N; i++ {
				ls := New(less)
				for j := 0; j < len(values); j++ {
					ls.Append(values[j])
				}
			}
		},
	)

	return r
}

func benchmarkGoList(b *testing.B, values ...interface{}) bool {
	r := b.Run(
		fmt.Sprintf("New Go list of %d values", len(values)),
		func(b0 *testing.B) {
			for i := 0; i < b0.N; i++ {
				ls := golist.New()
				for j := 0; j < len(values); j++ {
					ls.PushBack(values[j])
				}
			}
		},
	)

	return r
}

func benchmarkSliceCopy(b *testing.B, values ...interface{}) bool {
	r := b.Run(
		fmt.Sprintf("Copied slice of %d values", len(values)),
		func(b0 *testing.B) {
			for i := 0; i < b0.N; i++ {
				lst := make([]interface{}, len(values))
				copy(lst, values)
			}
		},
	)

	return r
}

func benchmarkSliceAppend(b *testing.B, values ...interface{}) bool {
	r := b.Run(
		fmt.Sprintf("Appended slice of %d values", len(values)),
		func(b0 *testing.B) {
			for i := 0; i < b0.N; i++ {
				lst := make([]interface{}, 0)
				for j := 0; j < len(values); j++ {
					lst = append(lst, values[j])
				}
			}
		},
	)

	return r
}
