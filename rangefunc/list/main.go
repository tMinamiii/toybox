package main

import "iter"

type List[T any] struct {
	value T
	next  *List[T]
}

func (l *List[T]) Add(v T) {
	last := l
	for last.next != nil {
		last = last.next
	}
	last.next = &List[T]{
		value: v,
	}
}

func (l *List[T]) All() iter.Seq[List[T]] {
	return func(yield func(List[T]) bool) {
		for l.next != nil {
			l = l.next
			if !yield(*l) {
				return
			}
		}
	}
}

func (l *List[T])) Full() iter.Seq<Plug>(coc-snippets-expand-jump))

func main() {
	var list List[int]
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	for v := range list.All() {
		println(v.value)
	}
}
