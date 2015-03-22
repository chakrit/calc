package main

import "testing"
import a "github.com/stretchr/testify/assert"

func TestStack_Simple(t *testing.T) {
	t1, t2 := &token{}, &token{}
	s := newStack(4)
	s.push(t1)
	s.push(t2)
	a.Equal(t, t2, s.pop())
	a.Equal(t, t1, s.pop())
	a.Nil(t, s.pop())
}

func TestStack_Peeked(t *testing.T) {
	t1, t2 := &token{}, &token{}
	s := newStack(4)
	a.Nil(t, s.peek())
	s.push(t1)
	a.Equal(t, t1, s.peek())
	s.push(t2)
	a.Equal(t, t2, s.peek())
	a.Equal(t, t2, s.pop())
	a.Equal(t, t1, s.pop())
	a.Nil(t, s.pop())
}
