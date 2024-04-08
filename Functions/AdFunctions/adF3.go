package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ad Func")
	test := func(y int) int {
		fmt.Println("func in main()\n", y)
		return y - 8

	}
	
	fmt.Println(test(99))
}
