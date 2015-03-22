package main

import "testing"
import a "github.com/stretchr/testify/assert"

func TestQueue_Simple(t *testing.T) {
	t1, t2 := &token{}, &token{}
	q := newQueue(4)
	q.enqueue(t1)
	q.enqueue(t2)

	o1, o2, o3 := q.dequeue(), q.dequeue(), q.dequeue()
	a.Equal(t, t1, o1)
	a.Equal(t, t2, o2)
	a.Nil(t, o3)
}

func TestQueue_Peeked(t *testing.T) {
	t1, t2 := &token{}, &token{}
	q := newQueue(4)
	q.enqueue(t1)
	a.Equal(t, t1, q.peek())
	q.enqueue(t2)
	a.Equal(t, t1, q.peek())
	a.Equal(t, t1, q.dequeue())
	q.enqueue(t1)
	a.Equal(t, t2, q.peek())
	a.Equal(t, t2, q.dequeue())
	a.Equal(t, t1, q.peek())
	a.Equal(t, t1, q.dequeue())
	a.Nil(t, q.peek())
	a.Nil(t, q.dequeue())
}
