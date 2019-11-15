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

func TestList(t *testing.T) {
	initSeed()

	var (
		iters    = 100
		numItems = 100
	)

	for i := 0; i < iters; i++ {
		var (
			ls   = New()
			nums = rand.Perm(numItems)
		)

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

		for i := 0; i < numItems; i++ {
			switch {
			case nums[i] != s[i].(int):
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, i, nums[i], s[i].(int))
			case nums[i] != m[i].(int):
				t.Fatalf("\nseed: %d\nexpected m[%d] = %d\nreceived %d\n", seed, i, nums[i], m[i].(int))
			}
		}

		// Test rotate & swap
		nums = append(nums[1:], nums[0]) // rotate left
		ls.RotateLeft()
		s = ls.Slice()
		for i := 0; i < numItems; i++ {
			if nums[i] != s[i] {
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, i, nums[i], s[i].(int))
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

		i, j := rand.Intn(numItems), rand.Intn(numItems)
		nums[i], nums[j] = nums[j], nums[i]
		ls.Swap(i, j)
		for i := 0; i < numItems; i++ {
			if nums[i] != s[i] {
				t.Fatalf("\nseed: %d\nexpected s[%d] = %d\nreceived %d\n", seed, i, nums[i], s[i].(int))
			}
		}
	}
}
