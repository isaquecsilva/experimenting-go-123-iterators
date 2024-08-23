package main

import "iter"

type Callback[T any] = func(int, T) T

type Map[T any] struct {
	slice []T
}

func (m Map[T]) Iter(callback Callback[T]) iter.Seq2[int, T] {
	return iter.Seq2[int, T](func(yield func(int, T) bool) {

		for i := 0; i < len(m.slice); i++ {
			newElement := callback(i, m.slice[i])

			if !yield(i, newElement) {
				return
			}
		}
	})
}

func NewMap[T any](s []T) Map[T] {
	return Map[T] {
		slice: s,
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, n := range NewMap(numbers).Iter(func(index, element int) int {
		return element * element
	}) {
		println(n)
	}
}