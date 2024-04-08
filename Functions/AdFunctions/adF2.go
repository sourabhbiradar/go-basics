package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ad Func")
	test := func(y int) {
		fmt.Println("func in main()\n", y)

	}
	test(99)
}
