package list

import (
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

// perm ...
func perm(n int) ([]int, *List) {
	var (
		s  = rand.Perm(n)
		ls = New()
	)

	for _, v := range s {
		ls.Append(v)
	}

	return s, ls
}

func remove(value int, s []int) []int {
	for i := 0; i < len(s); i++ {
		if s[i] == value {
			s = append(s[:i], s[i+1:]...)
		}
	}

	return s
}

func removeAt(index int, s []int) []int {
	return append(s[:index], s[index+1:]...)
}

func TestSearchList(t *testing.T) {
	initSeed()

	var numItems = 5
	for iters := 0; iters < 10; iters++ {
		nums, ls := perm(numItems)
		for j, v := range nums {
			if index, ok := ls.Search(v); j != index || !ok {
				fmt.Printf("\n%d\n%s\n", nums, ls.String())
				t.Fatalf("\nseed: %d\nexpected (%d, %t)\nreceived (%d, %t)\n", seed, j, true, index, ok)
			}
		}
	}
}

func TestSliceMapList(t *testing.T) {
	initSeed()

	var numItems = 100
	for iters := 0; iters < 10; iters++ {
		var (
			nums, ls = perm(numItems)
			s, m     = ls.Slice(), ls.Map()
		)

		switch {
		case len(s) != numItems:
			t.Fatalf("\nseed: %d\nexpected length %d\nreceived %d\n", seed, numItems, len(s))
		case len(m) != numItems:
			t.Fatalf("\nseed: %d\nexpected length %d\nreceived %d\n", seed, numItems, len(m))
		}

		for j := 0; j < numItems; j++ {
			switch {
			case nums[j] != s[j].(int):
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, j, nums[j], s[j].(int))
			case nums[j] != m[j].(int):
				t.Fatalf("\nseed: %d\nexpected m[%d] = %d\nreceived %d\n", seed, j, nums[j], m[j].(int))
			}
		}
	}
}

func TestRotateSwap(t *testing.T) {
	initSeed()

	var numItems = 100
	for iters := 0; iters < 10; iters++ {
		nums, ls := perm(numItems)

		nums = append(nums[1:], nums[0]) // rotate left
		ls.RotateLeft()
		s := ls.Slice()
		for j := 0; j < numItems; j++ {
			if nums[j] != s[j] {
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, j, nums[j], s[j].(int))
			}
		}

		nums = append([]int{nums[numItems-1]}, nums[:numItems-1]...) // rotate right
		ls.RotateRight()
		s = ls.Slice()
		for i := 0; i < numItems; i++ {
			if nums[i] != s[i] {
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, i, nums[i], s[i].(int))
			}
		}

		a, b := rand.Intn(numItems), rand.Intn(numItems)
		nums[a], nums[b] = nums[b], nums[a]
		ls.Swap(a, b)
		s = ls.Slice()
		for j := 0; j < numItems; j++ {
			if nums[j] != s[j] {
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, j, nums[j], s[j].(int))
			}
		}
	}
}

func TestRemove(t *testing.T) {
	initSeed()

	numItems := 100
	for i := 0; i < 10; i++ {
		var (
			nums, ls = perm(numItems)
			d        = nums[0]
		)

		// Remove values
		nums = remove(d, nums)
		ls.Remove(d)
		s := ls.Slice()
		for j := 0; j < len(nums); j++ {
			if nums[j] != s[j] {
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, j, nums[j], s[j].(int))
			}
		}

		// Remove index
		index := rand.Intn(len(nums))
		nums = removeAt(index, nums)
		ls.RemoveAt(index)
		s = ls.Slice()
		for j := 0; j < len(nums); j++ {
			if nums[j] != s[j] {
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, j, nums[j], s[j].(int))
			}
		}
	}
}
