package main

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

func isOp(r rune) bool {
	switch r {
	case '+', '-', '*', '/':
		return true
	}

	return false
}

func isWhitespace(r rune) bool {
	switch r {
	case ' ', '\r', '\n':
		return true
	}

	return false
}
