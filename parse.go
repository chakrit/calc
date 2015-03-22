package main

import "fmt"
import "strconv"

func parse(tokens <-chan *token) node {
	return parseRoot(tokens)
}

func parseRoot(tokens <-chan *token) *rootNode {
	return &rootNode{parseArithmetic(tokens)}
}

func parseArithmetic(tokens <-chan *token) node {
	operators := newStack(32)
	output := newStack(64)

	p := map[string]int{
		"+": 10, "-": 10, "*": 20, "/": 20, "(": 0, ")": 0,
	}

	for token := range tokens {
		switch token.tokenType {
		case typeNum:
			output.push(token)
		case typeLParen:
			operators.push(token)

		case typeOp:
			op := operators.peek()
			for op != nil && op.tokenType != typeLParen && p[token.text] < p[op.text] {
				output.push(operators.pop())
				op = operators.peek()
			}

			operators.push(token)

		case typeRParen: // TODO: Check paren balancing.
			for op := operators.pop(); op.tokenType != typeLParen; op = operators.pop() {
				output.push(op)
			}

		default:
			panic("invalid arithmetic.")
		}

		if DEBUG {
			fmt.Println(output, operators)
		}
	}

	for operators.peek() != nil {
		output.push(operators.pop())
	}

	var build func() node
	build = func() node {
		t := output.pop()
		switch t.tokenType {
		case typeNum:
			return &numberNode{mustAtoi(t.text)}
		case typeOp:
			rchild, lchild := build(), build()
			return &arithmeticNode{lchild, t.text, rchild}
		default:
			panic("invalid arithmetic.")
		}
	}

	return build()
}

func mustAtoi(src string) int {
	if n, e := strconv.Atoi(src); e != nil {
		panic(e)
	} else {
		return n
	}
}
