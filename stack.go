package main

import "fmt"
import "strings"

type stack struct {
	items []*token
}

func newStack(capacity int) *stack {
	return &stack{make([]*token, 0, capacity)}
}

func (s *stack) peek() *token {
	if len(s.items) == 0 {
		return nil
	}

	return s.items[len(s.items)-1]
}

func (s *stack) push(t *token) {
	s.items = append(s.items, t)
}

func (s *stack) pop() *token {
	if len(s.items) == 0 {
		return nil
	}

	idx := len(s.items) - 1
	result := s.items[idx]
	s.items = s.items[:idx]
	return result
}

func (s *stack) String() string {
	strs := make([]string, len(s.items))
	for i, item := range s.items {
		strs[i] = fmt.Sprint(item)
	}

	return "stack{" + strings.Join(strs, ", ") + "}"
}
