package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ad Func")
	test := func(y int) int {

		return y * -1

	}(8) // parameter //same as test(99)

	fmt.Println(test)
}
