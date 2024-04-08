package main

import (
	"fmt"
)

func roll(r int) int {
	if r == 0 {
		return 1
	}
	//return r
	return r * roll(r-1)//5!

}

func main() {
	fmt.Println("Recursion/Recursive functions")
	fmt.Println(roll(5)) //0 or string

	var clsr func(n int) int
	clsr = func(n int) int {
		if n < 7 {
			return n
		}
		return clsr(n - 1) //+ clsr(n-2)
	}
	fmt.Println(clsr(7))
}
