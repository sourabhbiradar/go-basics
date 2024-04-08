package main

import (
	"fmt"
)

func main() {
	fmt.Println("bool")

	x := []int{3, 4, 5, 6, 7}
	y := x

	fmt.Println(x[len(x)-1] == 7)
	fmt.Println(x[len(x)-1] == 6)
	fmt.Println(x[3] == y[3] && y[3] == 6)

}
