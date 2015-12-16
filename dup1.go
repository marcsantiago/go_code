package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) //dictionary here the key is of type string and the value is of type value
	input := bufio.NewScanner(os.Stdin)
	// this loop breaks when it reaches EOF
	for input.Scan() {
		// the key is equal to the text entered from the input stream and the value is set to 1
		// if the the key exist already the value is incremented.
		counts[input.Text()]++
	}
	//Note ignoring potential errors from input.Err()
	for line, n := range counts {
		// if the line was duplicated then print it out
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
