package main

import "fmt"

func parseMain(line string) string {
	context := newContext(line)

	go func() {
		defer context.done()
		f := parseFunc(parseStart)
		for f != nil {
			f = f(context)
		}
	}()

	result := compile(context.tokens)
	fmt.Println()
	return result
}

type parseFunc func(c *context) parseFunc

func parseStart(c *context) parseFunc {
	switch classOf(c.peek()) {
	case classNum:
		return parseNum
	case classOp:
		return parseOp
	case classWhite:
		c.consume()
		return parseStart

		// TODO: Parser error
	case classEOF:
		return nil
	default:
		return nil
	}
}

func parseNum(c *context) parseFunc {
	numStr := ""
	for r := c.consume(); classOf(r) == classNum; r = c.consume() {
		numStr += string(r)
	}

	c.backtrack() // last token is non-num
	c.emit(&token{typeNum, numStr})
	return parseStart
}

func parseOp(c *context) parseFunc {
	if r := c.consume(); classOf(r) == classOp {
		c.emit(&token{typeOp, string(r)})
	} else {
		c.backtrack()
		// TODO: ERROR
	}

	return parseStart
}
