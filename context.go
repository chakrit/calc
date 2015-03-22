package main

import "io"
import "bufio"

type context struct {
	scanner io.RuneScanner
	tokens  chan *token
	isEOF   bool
}

func newContext(reader io.Reader) *context {
	c := &context{
		scanner: bufio.NewReader(reader),
		tokens:  make(chan *token, 32),
		isEOF:   false,
	}
	return c
}

func (c *context) emit(t *token) {
	c.tokens <- t
}

func (c *context) eof() bool {
	return c.isEOF
}

func (c *context) done() {
	close(c.tokens)
}

func (c *context) peek() rune {
	r := c.consume()
	c.backtrack()
	return r
}

func (c *context) consume() rune {
	r, _, e := c.scanner.ReadRune()
	if e == io.EOF {
		c.isEOF = true
		return rune(0)
	} else {
		noe(e)
	}

	return r
}

func (c *context) backtrack() {
	c.scanner.UnreadRune()
}
