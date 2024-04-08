package main

import (
	"fmt"
)

func vf(h ...int) { //multiple arguement in func'(...)'
	//fmt.Println(h, "")
	total := 0
	for _, h := range h {
		total += h
	}
	fmt.Println(h, total)

}

func main() {
	fmt.Println("Varaidic Functions")
	vf(2, 3, 4, 5)
	vf(2, 3, 4, 5, 6)
	sums := []int{5, 6, 7, 8}
	vf(sums...)
}
