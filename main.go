package main

import (
	"fmt"
	"iter"
	"slices"
)

type Iterator[V any] struct {
	seq iter.Seq[V]
}

func From[V any](slice []V) *Iterator[V] {
	return &Iterator[V]{
		seq: slices.Values(slice),
	}
}

func (it *Iterator[V]) Filter(pred func(V) bool) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for item := range it.seq {
				if !pred(item) {
					continue
				}

				if !yield(item) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Map(mapper func(V) V) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for item := range it.seq {
				mapped := mapper(item)

				if !yield(mapped) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Reverse() *Iterator[V] {
	collected := it.Collect()
	slices.Reverse(collected)
	return &Iterator[V]{
		seq: slices.Values(collected),
	}
}

func (it *Iterator[V]) Take(quantity int) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			count := 0
			for item := range it.seq {
				if count >= quantity {
					return
				}

				if !yield(item) {
					return
				}

				count++
			}
		},
	}
}

func (it *Iterator[V]) Skip(quantity int) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			count := 0
			for item := range it.seq {
				if count < quantity {
					count++
					continue
				}

				if !yield(item) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Collect() []V {
	return slices.Collect(it.seq)
}

func main() {
	result := From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(x int) bool { return x%2 == 0 }).
		Map(func(x int) int { return x * 10 }).
		Collect()

	fmt.Println(result)

	resultReversed := From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(x int) bool { return x%2 == 0 }).
		Map(func(x int) int { return x * 10 }).
		Reverse().
		Collect()

	fmt.Println(resultReversed)

	resultLimited := From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(x int) bool { return x%2 == 0 }).
		Map(func(x int) int { return x * 10 }).
		Reverse().
		Take(2).
		Collect()

	fmt.Println(resultLimited)

	resultSkipped := From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(x int) bool { return x%2 == 0 }).
		Map(func(x int) int { return x * 10 }).
		Reverse().
		Skip(2).
		Take(2).
		Collect()

	fmt.Println(resultSkipped)
}
