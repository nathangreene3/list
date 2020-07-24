package list

import (
	"container/heap"
	golist "container/list"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"

	"github.com/nathangreene3/math"
)

// testInt is a comparable integer.
type testInt int

// Compare two integers.
func (a testInt) Compare(b Comparable) int {
	switch {
	case a < b.(testInt):
		return -1
	case b.(testInt) < a:
		return 1
	default:
		return 0
	}
}

// TestList ensures manipulating a list is equivalent to manipulating a slice.
func TestList(t *testing.T) {
	var (
		numTests = 8
		numItems = 256
	)

	for i := 0; i < numTests; i++ {
		var (
			ls   = New(Ints)
			nums = make([]int, 0, numItems)
		)

		for j := 0; j < numItems; j++ {
			x := rand.Int()
			ls.Append(x)
			nums = append(nums, x)
		}

		for j, itm := 0, ls.head; j < numItems && itm != nil; j, itm = j+1, itm.next {
			if index, ok := ls.Search(nums[j]); j != index || !ok {
				t.Fatalf("\nexpected (%d, %t)\nreceived (%d, %t)\n", j, true, index, ok)
			}
		}

		if s := ls.Slice(); len(s) != ls.length {
			t.Fatalf("\nexpected length %d\nreceived %d\n", ls.length, len(s))
		}

		if s := ls.Sort(); !sort.IsSorted(s) {
			t.Fatalf("\nexpected %v to be sorted\n", s)
		}
	}
}

func TestComparable(t *testing.T) {
	ls := New(CmpLess)
	for i := 0; i < 8; i++ {
		ls.Append(testInt(rand.Intn(10)))
	}

	ls.Sort()
	if !sort.IsSorted(ls) {
		t.Fatalf("\nexpected sorted list\nreceived %v\n", ls)
	}
}

// TestInsertRemove tests the manual alteration of a list's state.
func TestInsertRemove(t *testing.T) {
	ls := New(Ints, 0, 1, 0, 2, 0)
	ls.RemoveAt(4)
	ls.RemoveAt(2)
	ls.RemoveAt(0)
	ls.Append(0, 4, 0, 5, 0).Remove(0).InsertAt(2, 3).Prepend(0).Append(6)

	exp := []int{0, 1, 2, 3, 4, 5, 6}
	if len(exp) != ls.Len() {
		t.Fatalf("\nexpected %v\nreceived %v\n", exp, ls.String())
	}

	for i := 0; i < len(exp); i++ {
		if index, ok := ls.Search(exp[i]); i != index || !ok {
			t.Fatalf("\nexpected %v\nreceived %v\n", exp, ls.String())
		}
	}
}

// TestHeap heapifies a list and sorts integers.
func TestHeap(t *testing.T) {
	var (
		rnd = []int{9, 0, 8, 1, 7, 2, 6, 3, 5, 4}
		exp = make([]int, len(rnd))
		ls  = New(Ints)
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

// TestMapFilterReduce generates a list of integers as {1, 2, ..., n}, then reduces the list to compute 1+2+...+n = (n^2+n)/2.
func TestMapFilterReduce(t *testing.T) {
	for n := 1; n <= 256; n <<= 1 {
		// Computing sums
		var (
			exp int         = n * (n + 1) / 2                                                 // 1+2+...+n
			gen Generator   = func(i int) interface{} { return i + 1 }                        // Generates {1, 2, 3, ..., 256}
			red Reducer     = func(x, y interface{}) interface{} { return x.(int) + y.(int) } //
			rec interface{} = Generate(n, gen, Ints).Reduce(red)                              // Returns interface{}, not int on purpose
		)

		if exp != rec {
			t.Fatalf("\nexpected %d\nreceived %d\n", exp, rec)
		}
	}

	{
		// Filter primes
		// TODO: Set exp with math.primes
		exp := append(make([]int, 0, 256), 2)
		for n := 3; n < 256; n += 2 {
			if math.IsPrime(n) {
				exp = append(exp, n)
			}
		}

		var (
			gen  Generator     = func(i int) interface{} { return i + 1 }                  // Generate {1, 2, 3, ..., 256}
			fltr Filterer      = func(x interface{}) bool { return math.IsPrime(x.(int)) } // Filter primes
			rec  []interface{} = Generate(256, gen, Ints).Filter(fltr).Slice()
		)

		if len(exp) != len(rec) {
			t.Fatalf("\nexpected %d\nreceived %d\n", exp, rec)
		}

		for i := 0; i < len(rec); i++ {
			if exp[i] != rec[i] {
				t.Fatalf("\nexpected %d\nreceived %d\n", exp, rec)
			}
		}
	}

	{
		// Joining directories into a path
		var (
			values []string    = []string{"a", "b", "c", "d", "e"}
			exp    string      = strings.Join(values, "/")
			gen    Generator   = func(i int) interface{} { return string('a' + byte(i)) }
			red    Reducer     = func(x, y interface{}) interface{} { return x.(string) + "/" + y.(string) }
			rec    interface{} = Generate(len(values), gen, Strings).Reduce(red)
		)

		if exp != rec {
			t.Fatalf("\nexpected %q\nreceived %q\n", exp, rec)
		}
	}
}

func BenchmarkList(b *testing.B) {
	var (
		n  int           = 256
		s  int           = 8
		vs []interface{} = make([]interface{}, n)
		f  Lesser        = Ints
	)

	{
		// Linear benchmarks
		for i := 0; i <= len(vs); i += s {
			benchmarkList(b, f, vs[:i]...)
		}

		for i := 0; i <= len(vs); i += s {
			benchmarkGoList(b, vs[:i]...)
		}

		for i := 0; i <= len(vs); i += s {
			benchmarkSliceCopy(b, vs[:i]...)
		}

		for i := 0; i <= len(vs); i += s {
			benchmarkSliceAppend(b, vs[:i]...)
		}
	}

	{
		// Exponential benchmark
		for i := 1; i <= n; i <<= 1 {
			benchmarkList(b, f, vs[:i]...)
		}

		for i := 1; i <= n; i <<= 1 {
			benchmarkGoList(b, vs[:i]...)
		}

		for i := 1; i <= n; i <<= 1 {
			benchmarkSliceCopy(b, vs[:i]...)
		}

		for i := 1; i <= n; i <<= 1 {
			benchmarkSliceAppend(b, vs[:i]...)
		}
	}
}

func benchmarkList(b *testing.B, less Lesser, values ...interface{}) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			ls := New(less)
			for j := 0; j < len(values); j++ {
				ls.Append(values[j])
			}
		}
	}

	return b.Run(fmt.Sprintf("New list of %d values", len(values)), f)
}

func benchmarkGoList(b *testing.B, values ...interface{}) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			ls := golist.New()
			for j := 0; j < len(values); j++ {
				ls.PushBack(values[j])
			}
		}
	}

	return b.Run(fmt.Sprintf("New Go list of %d values", len(values)), f)
}

func benchmarkSliceCopy(b *testing.B, values ...interface{}) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			lst := make([]interface{}, len(values))
			copy(lst, values)
		}
	}

	return b.Run(fmt.Sprintf("Copied slice of %d values", len(values)), f)
}

func benchmarkSliceAppend(b *testing.B, values ...interface{}) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			lst := make([]interface{}, 0)
			for j := 0; j < len(values); j++ {
				lst = append(lst, values[j])
			}
		}
	}

	return b.Run(fmt.Sprintf("Appended slice of %d values", len(values)), f)
}
