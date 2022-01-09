package set

import (
	"fmt"
	"strings"
)

var empty struct{}

type Set[T comparable] map[T]struct{}

func (s *Set[T]) Add(v T) {
	(*s)[v] = empty
}

func (s *Set[T]) Remove(v T) {
	delete(*s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) String() string {
	var sb strings.Builder
	sb.WriteRune('{')

	first := true
	for k := range s {
		if !first {
			sb.WriteString(", ")
		}
		fmt.Fprint(&sb, k)
		first = false
	}

	sb.WriteRune('}')
	return sb.String()
}
