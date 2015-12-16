package main

import (
	"fmt"
)

func main() {
	var balance = 100.00
	var fee = 0.50
	// neat way of shorting the print function
	var print = fmt.Println
	var exit = true
	var widthdrawl int
	for {
		fmt.Scan(&widthdrawl)
		if widthdrawl%20 != 0 {
			print("Please try again.  Enter a multiple of 20")
			exit = false
		} else if float64(widthdrawl)+fee > balance {
			print("Error, widthdrawl amoun exceeds your balance of %s", balance)
			exit = false
		} else {
			exit = true
		}
		if exit {
			print(exit)
			break
		}
	}
	print("Remaining balance = %s", balance-(float64(widthdrawl)+fee))
}
