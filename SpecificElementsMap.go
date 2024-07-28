package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := a["brand"]
	val2, ok2 := a["color"]
	val3, ok3 := a["day"]
	_, ok4 := a["color"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
}
