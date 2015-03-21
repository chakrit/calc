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
		msg = "NUM"
	case typeOp:
		msg = "OP "
	case typeLeftParen:
		msg = "P( "
	case typeRightParen:
		msg = "P) "
	}

	msg += " " + t.text
	return msg
}
