package main

import "fmt"
import "os"

func main() {
	lines := make(chan string)
	results := parse(lines)

	go func() {
		var e error
		for line := "0"; line != "q"; _, e = fmt.Scanln(&line) {
			noe(e)
			lines <- line
		}
	}()

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("done.")
}

func parse(lines <-chan string) <-chan string {
	result := make(chan string)

	go func() {
		defer close(result)
		for line := range lines {
			result <- parseMain(line)
		}
	}()

	return result
}

func noe(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1) //noreturn
	}
}
