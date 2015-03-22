package main

import "fmt"
import "os"

const DEBUG = false

func main() {
	result := compile(parse(lex(os.Stdin)))
	fmt.Fprintln(os.Stdout, result)
}

func noe(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1) //noreturn
	}
}
