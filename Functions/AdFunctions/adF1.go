package main

import (
	"fmt"
)

func test(y int) {
	fmt.Println("hello", y)
}

func main() {
	fmt.Println("Advanced Functions")
	x := test
	//x()
	x(5)
}
