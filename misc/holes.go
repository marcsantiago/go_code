package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var holes = []string{"o", "R", "g", "p", "P", "q", "Q", "d", "D", "b", "B"}
	var count = 0
	str = "Test String"

	for i := range holes {
		if strings.Contains(str, holes[i]) {
			count++
		}
	}
	fmt.Println(count)
}
