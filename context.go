package main

import "io"
import "strings"

type context struct {
	*strings.Reader
	tokens chan *token
	result int
	isEOF  bool
}

func newContext(line string) *context {
	c := &context{
		Reader: strings.NewReader(line),
		tokens: make(chan *token, 32),
		result: 0,
		isEOF:  false,
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
	r, _, e := c.Reader.ReadRune()
	if e == io.EOF {
		c.isEOF = true
		return rune(0)
	} else {
		noe(e)
	}

	return r
}

func (c *context) backtrack() {
	c.Reader.UnreadRune()
}
