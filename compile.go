package main

import "fmt"
import "strconv"

func compile(root node) string {
	if DEBUG {
		fmt.Println(root)
	}

	n := compileArithmetic(root.(*rootNode).root)
	return strconv.Itoa(n)
}

func compileArithmetic(node interface{}) int {
	switch n := node.(type) {
	case *numberNode:
		return n.n
	case *arithmeticNode:
		l, r := compileArithmetic(n.lchild), compileArithmetic(n.rchild)
		switch n.op {
		case "+":
			return l + r
		case "-":
			return l - r
		case "*":
			return l * r
		case "/":
			return l / r
		}
	}

	panic("bad parse.")
}
