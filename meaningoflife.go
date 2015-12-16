package main

import (
	"fmt"
)

// func main() {
// 	life := 42
// 	var x int
// 	for {
// 		fmt.Scan(&x)
// 		if x == life {
// 			fmt.Println("Correct the meaning of life is 42")
// 			break
// 		}
// 	}
// }

// func main() {
//   for {
//     var x int
//     fmt.Scan(&x)
//     if x == 42 {
//       fmt.Println("Correct the meaning of life is 42")
//       break
//     }
//   }
// }

func main() {
	for {
		var x int
		fmt.Scan(&x)
		if x == 42 {
			break
		} else {
			fmt.Println(x)
		}
	}
}
