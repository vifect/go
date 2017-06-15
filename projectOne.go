package main

import (
	"fmt"
)

func add(x int, y int) [3]int {
	var a [3]int
	a[0] = x
	a[1] = y
	a[2] = x + y
	return a
}

func main() {

	fmt.Println(add(42, 43))
}
