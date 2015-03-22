package main

func lexMain(line string) chan *token {
	context := newContext(line)

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
	switch classOf(c.peek()) {
	case classNum:
		return lexNum
	case classOp:
		return lexOp
	case classWhite:
		c.consume()
		return lexStart

		// TODO: lexer error
	case classEOF:
		return nil
	default:
		return nil
	}
}

func lexNum(c *context) lexFunc {
	numStr := ""
	for r := c.consume(); classOf(r) == classNum; r = c.consume() {
		numStr += string(r)
	}

	c.backtrack() // last token is non-num
	c.emit(&token{typeNum, numStr})
	return lexStart
}

func lexOp(c *context) lexFunc {
	if r := c.consume(); classOf(r) == classOp {
		c.emit(&token{typeOp, string(r)})
	} else {
		c.backtrack()
		// TODO: ERROR
	}

	return lexStart
}
