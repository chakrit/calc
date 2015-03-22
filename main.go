package main

import "fmt"
import "os"
import "bufio"

func main() {
	lines := make(chan string)
	results := parse(lines)

	go func() {
		defer close(lines)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines <- scanner.Text()
			noe(scanner.Err())
		}
	}()

	for result := range results {
		fmt.Println()
		fmt.Println("result = ", result)
	}

	fmt.Println("done.")
}

func parse(lines <-chan string) <-chan string {
	result := make(chan string)

	go func() {
		defer close(result)
		for line := range lines {
			result <- compile(lexMain(line))
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
