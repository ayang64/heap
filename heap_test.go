package heap

import (
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	cmp := func(a, b int) bool {
		return a < b
	}

	h := New(cmp)

	t.Logf("pushing")

	for range 15 {
		h.Push(rand.Intn(100))
	}

	t.Logf("-> %#v", h.s)

	t.Logf("popping")
	for v := range h.All() {
		t.Logf("%d", v)
	}
}
