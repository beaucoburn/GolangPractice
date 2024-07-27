package main

import (
	"fmt"
)

func myFunction(x int, y int) (result int) {
	result = x + y
	return result
}

func main() {
	fmt.Println(myFunction(1, 2))
}
