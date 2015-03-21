package main

import "fmt"
import "strconv"

func compile(tokens <-chan *token) string {
	n, op := 0, ""

	for token := range tokens {
		fmt.Print(token, " ")

		switch token.tokenType {
		case typeNum:
			num, e := strconv.Atoi(token.text)
			noe(e)

			if op != "" {
				switch op {
				case "+":
					num = n + num
				case "-":
					num = n - num
				case "*":
					num = n * num
				case "/":
					num = n / num
				}
				op = ""
			}

			n = num
		case typeOp:
			op = token.text
		}
	}

	return strconv.Itoa(n)
}
