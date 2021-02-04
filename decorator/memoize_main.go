package main

import (
	"fmt"
	"go-design-patterns/decorator/memoize"
)

func main() {

	decorated := memoize.Memoize(func(x int) int {
		return x + 1
	})
	x := decorated(1)
	fmt.Println(x)
	y := decorated(1)
	fmt.Println(y)
}
