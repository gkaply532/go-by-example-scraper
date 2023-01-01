package set

type Set[T comparable] struct {
	elements map[T]*any
}

func New[T comparable]() Set[T] {
	return Set[T]{elements: make(map[T]*any)}
}

func (s Set[T]) Add(v T) {
	s.elements[v] = nil
}

func (s Set[T]) First() (T, bool) {
	for k := range s.elements {
		return k, true
	}

	var noop T
	return noop, false
}

func (s Set[T]) Remove(v T) {
	delete(s.elements, v)
}

func (s Set[T]) Present(v T) bool {
	_, pres := s.elements[v]
	return pres
}

func (s Set[T]) Len() int {
	return len(s.elements)
}
