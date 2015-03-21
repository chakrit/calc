package main

type class int

const (
	classNum class = iota
	classOp
	classWhite
	classEOF
)

func classOf(r rune) class {
	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return classNum
	case '+', '-', '*', '/':
		return classOp
	case ' ', '\r', '\n':
		return classWhite
	case rune(0):
		return classEOF
	default:
		// ERROR!
		return classEOF
	}
}
