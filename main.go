package main

import "fmt"
import "os"

func main() {
	results := compile(lex(os.Stdin))
	for result := range results {
		fmt.Fprintln(os.Stdout, result)
	}
}

func noe(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1) //noreturn
	}
}
