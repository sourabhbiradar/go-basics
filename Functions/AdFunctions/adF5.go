package main

import (
	"fmt"
)

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(9))
}

func main() {
	fmt.Println("Ad Func")
	test := func(y int) int {

		return y * -1

	}

	test2(test)
}
