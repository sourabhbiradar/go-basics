package main

import (
	"fmt"
)

func main() {
	fmt.Println("slice")
	s := make([]string, 4)
	fmt.Println(s)
	s[0] = "emp1"
	s[1] = "emp2"
	s[2] = "emp3"
	fmt.Println("set :", s)
	fmt.Println("get :", s[1])
	fmt.Println("length :", len(s))

	s = append(s, "emp4") //append
	s = append(s, "emp5", "emp6")
	fmt.Println("appended set :", s)

	c := make([]string, len(s))
	copy(c, s) //copy
	fmt.Println("copied slice from 's' :", c)

	fmt.Println("slice operator with syntax 'slice[low:high]'")
	l := s[1:5] //shud be emp1,emp2,...,emp4 /s[1] to s[4]
	fmt.Println(l)
	l = s[:6]
	fmt.Println(l)

	fmt.Println("single dimensional slice")
	oneD := []string{"A", "B", "C"}
	fmt.Println("1D :", oneD)

	fmt.Println("two dimensional slice")
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(twoD)

	twoDi := make([][]int, 3) //2D slice using make()
	twoDi[0] = []int{1, 2, 3}
	twoDi[1] = []int{4, 5, 6}
	//twoDi[2][0] = 7 //element by element nt possible
	//twoDi[2][1]=8
	//twoDi[2][2]=9
	fmt.Println(twoDi)

}
