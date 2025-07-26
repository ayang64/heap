package heap

import (
	"fmt"
	"iter"
)

type Heap[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func NewWithCap[T any](cmp func(T, T) bool, c int) *Heap[T] {
	return &Heap[T]{cmp: cmp, s: make([]T, 0, c)}
}

func New[T any](cmp func(T, T) bool) *Heap[T] {
	return &Heap[T]{cmp: cmp}
}

func dup[T any](s []T) []T {
	r := make([]T, len(s))
	copy(r, s)
	return r
}

func (h *Heap[T]) Clone() *Heap[T] {
	return &Heap[T]{cmp: h.cmp, s: dup(h.s)}
}

func (h *Heap[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			v, err := h.Pop()
			if err != nil || !yield(v) {
				return
			}
		}
	}
}

func (h *Heap[T]) Empty() bool {
	return len(h.s) == 0
}

func (h *Heap[T]) PeekOr(v T) T {
	if len(h.s) == 0 {
		return v
	}
	return h.s[0]
}

func (h *Heap[T]) Peek() (T, error) {
	if len(h.s) == 0 {
		var zero T
		return zero, fmt.Errorf("empty heap")
	}
	return h.s[0], nil
}

func (h *Heap[T]) Len() int {
	return len(h.s)
}

func (h *Heap[T]) Push(v T) {
	h.s = append(h.s, v)
	up(h.cmp, h.s)
}

func (h *Heap[T]) Pop() (T, error) {
	if len(h.s) == 0 {
		var zero T
		return zero, fmt.Errorf("no more entries of type %T", zero)
	}
	v := h.s[0]
	h.s[0] = h.s[len(h.s)-1]
	h.s = h.s[:len(h.s)-1]

	down(h.cmp, h.s)
	return v, nil
}

func up[T any](cmp func(T, T) bool, s []T) {
	for cur := len(s) - 1; cur > 0; {
		switch parent := (cur - 1) / 2; cmp(s[cur], s[parent]) {
		case true:
			s[cur], s[parent] = s[parent], s[cur]
			cur = parent
		default:
			return
		}
	}
}

func down[T any](cmp func(T, T) bool, s []T) {
	for cur := 0; cur < len(s); {
		left := cur*2 + 1
		right := left + 1
		lowest := cur
		if right < len(s) && cmp(s[right], s[lowest]) {
			lowest = right
		}
		if left < len(s) && cmp(s[left], s[lowest]) {
			lowest = left
		}
		if lowest == cur {
			return
		}
		s[lowest], s[cur] = s[cur], s[lowest]
		cur = lowest
	}
}
