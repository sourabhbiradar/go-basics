package main

import (
	"fmt"
	"time"
)

func Switch() {
	t := time.Now()
	fmt.Println(t)
	k := 1
	fmt.Print(k)
	switch k {
	case 0:
		fmt.Println(" is zero")
	case 2:
		fmt.Println(" is even ")
	case 1: //k='1' so case '1' executed
		fmt.Println(" is odd")
	}
	p := "Sachin"
	fmt.Print(p)
	switch p {
	case "VK":
		fmt.Println("17")
	case "LM", "Sachin": //executes if even one is true
		fmt.Println("10")
	case "CR": //p="CR" so case "CR" executed
		fmt.Println("7")

	}

}

func dflt(q int) {
	q = 0
	switch {

	case q%2 != 0:
		fmt.Printf("%v is odd number", q)
	default:
		fmt.Printf("%v is even number", q)

	}
}

func main() {
	fmt.Println("Switch")
	Switch()
	dflt(3)
}
