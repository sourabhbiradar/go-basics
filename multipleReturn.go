package main

import (
	"fmt"
)

func num() (int, string, float32) {
	return 7, "seven", 66.98
}

func main() {

	fmt.Println("Multiple Return Values")

	a, b, c := num()
	fmt.Println(a, b, c)

	_, f, _ := num()
	fmt.Println(f)
}
