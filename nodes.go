package main

import "fmt"

type node interface {
}

type rootNode struct {
	root node
}

func (n *rootNode) String() string {
	return "root(" + fmt.Sprint(n.root) + ")"
}

type arithmeticNode struct {
	lchild node
	op     string
	rchild node
}

func (n *arithmeticNode) String() string {
	return n.op + "(" + fmt.Sprint(n.lchild) + ", " + fmt.Sprint(n.rchild) + ")"
}

type numberNode struct {
	n int
}

func (n *numberNode) String() string {
	return fmt.Sprint(n.n)
}
