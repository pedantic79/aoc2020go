package util

import (
	"fmt"
	"reflect"
	"strings"
)

var Empty struct{}

func setToString(set interface{}) string {
	var sb strings.Builder
	sb.WriteRune('{')

	first := true
	for _, k := range reflect.ValueOf(set).MapKeys() {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprint(k))
		first = false
	}

	sb.WriteRune('}')
	return sb.String()
}

type Set map[interface{}]struct{}

func NewSet() Set {
	return make(Set)
}

func (s *Set) Add(v interface{}) {
	(*s)[v] = Empty
}

func (s *Set) Remove(v interface{}) {
	delete(*s, v)
}

func (s Set) Contains(v interface{}) bool {
	_, ok := s[v]
	return ok
}

func (s Set) String() string {
	return setToString(s)
}
