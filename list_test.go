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

func TestList(t *testing.T) {
	initSeed()

	var (
		iters    = 100
		numItems = 100
	)

	for i := 0; i < iters; i++ {
		ls := New()
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
			t.Fatalf("\nexpected length %d, received %d\n", ls.length, len(s))
		}
	}
}
