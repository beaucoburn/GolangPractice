package main

import (
	"fmt"
)

func myFunction(x int, y int) (result int) {
	result = x + y
	return
}

func main() {
	total := myFunction(1, 2)
	fmt.Println(total)
}
