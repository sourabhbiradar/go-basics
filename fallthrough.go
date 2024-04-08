package main

import (
	"fmt"
)

func main() {
	fmt.Println("fallthrough")

	f := 12
	switch f {
	case 2:
		fmt.Println("2")
		fallthrough
	case 10:
		fmt.Println("10")
		fallthrough
	case 12:
		fmt.Println("12")
		fallthrough
	case 13:
		fmt.Println("13")
		fallthrough
		//fallthrough
	case 14:
		fmt.Println("14")
		fallthrough

	default:
		fmt.Println("default")

	}
}
