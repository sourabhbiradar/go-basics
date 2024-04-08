package main

import (
	"fmt"
)

func main() {
	fmt.Println("[...] array")

	x := [...]int{3, 4, 5, 6}
	fmt.Println(x, x[2])

	y := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(y[:4], y[3:])

	z := [...]int{10: -6}
	//{10th element is -6}
	fmt.Println(z)

	a := [7]int{4: 38}
	fmt.Println(a, "\n4th element :", a[4])
}
