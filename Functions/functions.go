package main

import (
	"fmt"
)

func add(x int, y int) int {

	return x + y

}

func multiply(a, b, c int) int {
	return a * b * c
}

func main() {
	fmt.Println("functions")

	sum := add(4, 7) //func add(x int , y int)int
	fmt.Println("sum :", sum)

	result := multiply(2, 3, 4)
	fmt.Println("result :", result)

}
