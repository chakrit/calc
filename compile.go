package main

import "strconv"

func compile(tokens <-chan *token) <-chan string {
	result := make(chan string)

	go func() {
		defer close(result)
		n, op := 0, ""

		for token := range tokens {
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

		result <- strconv.Itoa(n)
	}()

	return result
}
