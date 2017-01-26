package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(isPalidrome(str))
}

func isPalidrome(str string) bool {
	var revStr string
	for i := len(str) - 1; i >= 0; i-- {
		revStr += string(str[i])
	}
	if str == revStr {
		return true
	}
	return false
}

// func isPalidrome2(str string) bool {
// 	n := len(str)
// 	if n%2 == 0 {
//
// 		h1 := sort.Sort(str[n/2:])
// 		h2 := sort.Sort(str[:n/2])
// 		if strings.Join(h1, "") == strings.Join(h2, "") {
// 			return true
// 		}
//
// 	} else {
// 		fmt.Println("")
// 	}
//   return false
// }
