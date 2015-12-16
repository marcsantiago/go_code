package main

import (
	"fmt"
	"strings" //http://www.dotnetperls.com/replace-go
)

func main() {
	var dna string
	fmt.Scan(&dna)

	fmt.Println(reverseCompliment(dna))

}

func compliment(str string) string {
	//create a dictionary that has a key value that are both strings
	//var dnaMap = make(map[string]string)
	var dnaMap = map[string]string{
		"A": "T",
		"T": "A",
		"C": "G",
		"G": "C",
	}
	var comp string
	str = strings.ToUpper(str)
	for i := 0; i < len(str); i++ {
		comp += dnaMap[string(str[i])]
	}
	return comp
}

func rev(str string) string {
	var rev string
	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	return rev
}

func reverseCompliment(str string) string {
	return rev(compliment(str))
}
