package set

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

var empty struct{}

type Set[T comparable] map[T]struct{}

func New[T comparable]() map[T]struct{} {
	return make(map[T]struct{})
}

func Add[S ~map[T]struct{}, T comparable](s S, v T) {
	s[v] = empty
}

func Delete[S ~map[T]struct{}, T comparable](s S, v T) {
	delete(s, v)
}

func DeleteFunc[S ~map[T]struct{}, T comparable](s S, del func(T) bool) {
	maps.DeleteFunc(s, func(k T, v struct{}) bool {
		return del(k)
	})
}

func Contains[S ~map[T]struct{}, T comparable](s S, v T) bool {
	_, ok := s[v]
	return ok
}

func Equal[S1, S2 ~map[T]struct{}, T comparable](s1 S1, s2 S2) bool {
	// This doesn't work on go1.18beta2
	// return maps.Equal(s1, s2)

	if len(s1) != len(s2) {
		return false
	}
	for k, v1 := range s1 {
		if v2, ok := s2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
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
