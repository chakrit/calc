package main

type tokenType int

const (
	typeNum tokenType = iota
	typeOp
	typeLeftParen
	typeRightParen
)

type token struct {
	tokenType tokenType
	text      string
}

func (t *token) String() string {
	msg := ""
	switch t.tokenType {
	case typeNum:
		msg = "num"
	case typeOp:
		msg = "op"
	case typeLeftParen:
		msg = "lparen"
	case typeRightParen:
		msg = "rparen"
	}

	msg = msg + "(" + t.text + ")"
	return msg
}
