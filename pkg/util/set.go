package util

import (
	"iter"
	"maps"
	"slices"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](capacity uint, items ...T) Set[T] {
	capacity = Max(capacity, uint(len(items)))

	ret := make(Set[T], capacity)
	for _, x := range items {
		ret.Add(x)
	}
	return ret
}

func (s *Set[T]) Contains(item T) bool {
	_, exists := (*s)[item]
	return exists
}

func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		(*s)[item] = struct{}{}
	}
}

func (s *Set[T]) Delete(items ...T) {
	for _, item := range items {
		delete(*s, item)
	}
}

func (s *Set[T]) Items() []T {
	return slices.Collect(s.ItemsSeq())
}

func (s *Set[T]) ItemsSeq() iter.Seq[T] {
	return maps.Keys(*s)
}

func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	ret := make(Set[T], len(*s)+len(*s2))
	for k := range *s {
		ret.Add(k)
	}
	for k := range *s2 {
		ret.Add(k)
	}
	return &ret
}

func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	ret := make(Set[T], len(*s)+len(*s2))
	for k := range *s {
		if s2.Contains(k) {
			ret.Add(k)
		}
	}
	return &ret
}

func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	ret := make(Set[T], len(*s)+len(*s2))
	for k := range *s {
		if !s2.Contains(k) {
			ret.Add(k)
		}
	}

	return &ret
}
