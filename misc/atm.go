package main

import (
	"fmt"
)

func main() {
	balance := 100.00
	fee := 0.50
	exit := true
	var widthdrawl int
	for {
		if exit {
			break
		}
		fmt.Scan(&widthdrawl)
		if widthdrawl%20 != 0 {
			fmt.Println("Please try again.  Enter a multiple of 20")
			exit = false
		} else if float64(widthdrawl)+fee > balance {
			fmt.Printf("Error, widthdrawl amoun exceeds your balance of %f", balance)
			exit = false
		} else {
			exit = true
		}

	}
	fmt.Printf("Remaining balance = %f", balance-(float64(widthdrawl)+fee))
}
