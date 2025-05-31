package model

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(val T) {
	s[val] = struct{}{}
}

func (s Set[T]) Remove(val T) {
	delete(s, val)
}

func (s Set[T]) Contains(val T) bool {
	_, ok := s[val]
	return ok
}

func (s Set[T]) Size() int {
	return len(s)
}
