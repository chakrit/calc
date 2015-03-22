package main

import "io"
import "fmt"
import "os"

const DEBUG = false

func main() {
	run(os.Stdin, os.Stdout)
}

func run(input io.Reader, output io.Writer) {
	fmt.Fprintln(output, compile(parse(lex(input))))
}

func noe(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1) //noreturn
	}
}
