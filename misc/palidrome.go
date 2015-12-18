package main

import (
	"fmt"
	//"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(isPalidrome(str))
}

func isPalidrome(str string) bool {
	var rev_str string
	for i := len(str) - 1; i >= 0; i-- {
		rev_str += string(str[i])
	}
	if str == rev_str {
		return true
	}
	return false
}
