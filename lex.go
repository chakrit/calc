package main

import "io"

func lex(input io.Reader) chan *token {
	context := newContext(input)

	go func() {
		defer context.done()
		f := lexFunc(lexStart)
		for f != nil {
			f = f(context)
		}
	}()

	return context.tokens
}

type lexFunc func(c *context) lexFunc

func lexStart(c *context) lexFunc {
	r := c.peek()
	switch {
	case isWhitespace(r):
		c.consume()
	case isNumber(r):
		return lexNum
	case isOp(r):
		return lexOp
	case r == '(':
		c.consume()
		c.emit(&token{typeLParen, "("})
	case r == ')':
		c.consume()
		c.emit(&token{typeRParen, ")"})
	case r == rune(0):
		return nil
	}

	return lexStart
}

func lexNum(c *context) lexFunc {
	numStr := ""
	for r := c.consume(); isNumber(r); r = c.consume() {
		numStr += string(r)
	}

	c.backtrack() // last token is non-num
	c.emit(&token{typeNum, numStr})
	return lexStart
}

func lexOp(c *context) lexFunc {
	if r := c.consume(); isOp(r) {
		c.emit(&token{typeOp, string(r)})
	} else {
		c.backtrack()
		// TODO: ERROR
	}

	return lexStart
}
