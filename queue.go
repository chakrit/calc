package main

type queue struct {
	buffer chan *token
	peeked *token
}

func newQueue(capacity int) *queue {
	return &queue{
		buffer: make(chan *token, capacity),
		peeked:   nil,
	}
}

func (q *queue) peek() *token {
	if q.peeked != nil {
		return q.peeked
	}

	select {
	case t := <-q.buffer:
		q.peeked = t
		return t
	default:
		return nil
	}
}

func (q *queue) dequeue() *token {
	if q.peeked != nil {
		result := q.peeked
		q.peeked = nil
		return result
	}

	select {
	case t := <-q.buffer:
		return t
	default:
		return nil
	}
}

func (q *queue) enqueue(t *token) {
	q.buffer <- t
}
