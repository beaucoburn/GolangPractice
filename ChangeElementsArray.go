package main

import (
	"fmt"
)

func main() {
	prices := [3]int{10, 20, 30}
	fmt.Println(prices)
	prices[2] = 50
	fmt.Println(prices)
}
