package main

import (
	"fmt"
)

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "!")
}

func main() {
	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)
}
