package main

import "fmt"

var (
	primes []int
)

func init() {
	primes = sieveOfEratosthenes(100000000)
}

func binarySearch(arr []int, value int) int {
	startIndex := 0
	endIndex := len(arr) - 1
	lastValue := 0
	for startIndex <= endIndex {
		median := (startIndex + endIndex) / 2
		if arr[median] < value {
			startIndex = median + 1
			lastValue = median + 1
		} else {
			endIndex = median - 1
		}
	}
	if startIndex == len(arr) || arr[startIndex] != value {
		return lastValue
	}
	return startIndex
}

func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func filter(low, high int) []int {
	start := binarySearch(primes, low)
	end := binarySearch(primes, high)
	return primes[start:end]
}

func main() {
	var input int
	var res [][]int

	fmt.Scan(&input)
	for i := 0; i < input; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		tmp := []int{}
		for _, n := range filter(x, y) {
			tmp = append(tmp, n)
		}
		res = append(res, tmp)
	}

	for _, item := range res {
		for _, n := range item {
			fmt.Println(n)
		}
		fmt.Println("")
	}
}

// Another way to do it

// package main
//
// import "fmt"
//
// func generate(ch chan<- int) {
//   for i := 2; ; i++ {
//     ch <- i // Send 'i' to channel 'ch'.
//   }
// }
//
// func filter(in <-chan int, out chan<- int, prime int) {
//   for {
//     i := <-in // Receive value from 'in'.
//     if i%prime != 0 {
//       out <- i // Send 'i' to 'out'.
//     }
//   }
// }
//
// func main() {
//   var rep int
//   var high, low int
//   fmt.Scan(&rep)
//   vals := make([][]int, rep)
//   for i := 0; i < rep; i++ {
//     fmt.Scan(&low, &high)
//     vals[i] = []int{low, high}
//   }
//
//   for i := range vals {
//     low, high := vals[i][0], vals[i][1]
//     ch := make(chan int)
//     go generate(ch)
//     for i := low; i < high; i++ {
//       prime := <-ch
//       if prime >= low && prime <= high {
//         fmt.Println(prime)
//       }
//       if prime >= high {
//         break
//       }
//       ch1 := make(chan int)
//       go filter(ch, ch1, prime)
//       ch = ch1
//
//     }
//     println("")
//   }
// }
