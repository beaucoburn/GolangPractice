package main

import (
	"fmt"
)

func familyName(fname string) {
	fmt.Println("Hello", fname, "!")
}

func main() {
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")
}
