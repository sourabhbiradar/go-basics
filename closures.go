package main

import (
	"fmt"
)

func clsr() func() int {
	i := 0
	return func() int {
		i++
		return i // replace i -> any int
	}
}

func main() {
	fmt.Println("closures")

	newClsr := clsr()
	fmt.Println(newClsr())
	fmt.Println(newClsr())
	fmt.Println(newClsr())
	//fmt.Println(newClsr())

	newClsr2 := clsr()
	fmt.Println(newClsr2())

}
