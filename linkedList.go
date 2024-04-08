package main

import (
	"fmt"
)

type link struct {
	x int
	y []int
	z *link //impt
	w *link //impt
}

func main() {
	fmt.Println("linked list")
	var a, b, c link
	a = link{2, []int{3, 99, 87, 02}, &c, &b}
	b = link{4, []int{7, 8, 9, 0}, &b, &a}
	c = link{7, []int{1, 2, 3, 4}, &a, &b}
	fmt.Println(a.z.x)
	fmt.Println(b.w.y[2])
	fmt.Println(c)

}
