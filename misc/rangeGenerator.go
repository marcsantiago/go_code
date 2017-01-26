package main

import (
	"fmt"
)

// RangeLister ...
type RangeLister struct {
	list []int
}

func (r RangeLister) xrange(a int, b int) []int {
	for i := a; i <= b; i++ {
		r.list = append(r.list, i)
	}

	return r.list
}

func main() {
	print := fmt.Println
	rl := new(RangeLister)
	print(rl.xrange(1, 10))
}
