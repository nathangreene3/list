package list

import (
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

func TestSearchList(t *testing.T) {
	// initSeed()

	var numItems = 100
	for iters := 0; iters < 1; iters++ {
		var (
			ls   = New()
			nums = make([]int, 0, numItems)
		)

		for j := 0; j < numItems; j++ {
			nums = append(nums, rand.Intn(10))
		}

		for _, x := range nums {
			ls.Append(x)
		}

		// Test search
		itm := ls.head
		for j := 0; j < numItems && itm != nil && itm.next != nil; itm = itm.next {
			if index, ok := ls.Search(nums[j]); j != index || !ok {
				t.Fatalf("\nseed: %d\nexpected (%d, %t)\nreceived (%d, %t)\n", seed, j, true, index, ok)
			}

			j++
		}

		// Test slice & map
		s, m := ls.Slice(), ls.Map()
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

		// Test rotate & swap
		nums = append(nums[1:], nums[0]) // rotate left
		ls.RotateLeft()
		s = ls.Slice()
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
		for j := 0; j < numItems; j++ {
			if nums[j] != s[j] {
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, j, nums[j], s[j].(int))
			}
		}
	}
}

func TestSliceMapList(t *testing.T) {
	// initSeed()

	var numItems = 100
	for iters := 0; iters < 1; iters++ {
		var (
			ls   = New()
			nums = make([]int, 0, numItems)
		)

		for j := 0; j < numItems; j++ {
			nums = append(nums, rand.Intn(10))
		}

		for _, x := range nums {
			ls.Append(x)
		}

		// Test slice & map
		s, m := ls.Slice(), ls.Map()
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

		// Test rotate & swap
		nums = append(nums[1:], nums[0]) // rotate left
		ls.RotateLeft()
		s = ls.Slice()
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
		for j := 0; j < numItems; j++ {
			if nums[j] != s[j] {
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, j, nums[j], s[j].(int))
			}
		}
	}
}
