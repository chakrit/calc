package main

type tokenType int

const (
	typeNum tokenType = iota
	typeOp
	typeLParen
	typeRParen
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
	case typeLParen:
		msg = "lparen"
	case typeRParen:
		msg = "rparen"
	}

	msg = msg + "(" + t.text + ")"
	return msg
}
