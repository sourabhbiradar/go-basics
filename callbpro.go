package main

import (
	"fmt"
)

var a, b, c int

func idproof(f func()) {
	h := 0
	fmt.Println("Choose IDProof ")
	fmt.Println(" 0 : Adhar\n", "1 : Pan\n", "2 : Passport")
	fmt.Scanln(&h)

	if h == 0 {
		fmt.Println("Please enter Adhar number")
		fmt.Scanln(&a)
	} else if h == 1 {
		fmt.Println("Please enter Pan number")
		fmt.Scanln(&b)
	} else if h == 2 {
		fmt.Println("Please enter Passport number")
		fmt.Scanln(&c)
	} else {
		fmt.Println("Select valid option")
		f()
	}
}
func f() {
	var h int
	fmt.Scanln(&h)
	if h == 0 {
		fmt.Println("Please enter Adhar number")
		fmt.Scanln(&a)
	} else if h == 1 {
		fmt.Println("Please enter Pan number")
		fmt.Scanln(&b)
	} else if h == 2 {
		fmt.Println("Please enter Passport number")
		fmt.Scanln(&c)
	} else {
		idproof(f)

	}
}

func main() {
	fmt.Println("callBack")
	fmt.Println("ID Proof")
	idproof(f)

}
