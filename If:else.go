package main

import "fmt"

func IfnElse() {
	god := !true
	if god == true {
		fmt.Println("god exists")
	} else {
		fmt.Println("god doesnt exist")
	}
	
	num := 10
	if num < 9 {

		fmt.Println("num is greater than :", num)

	} else if num != 10 {
		fmt.Println("num is eqaul to :", num)
	} else {
		fmt.Println("num is lesser than :", num)
	}
}

func main() {
	IfnElse()
}
