package main

import (
	"fmt"
)

func main() {
	fmt.Println("Range")

	x := []int{1, 2, 3, 4}

	for i, v := range x {
		fmt.Println(i, v)
	}

	for _, x := range x {
		fmt.Println(x)
	}

	sum := 0
	for _, v := range x { //v==x
		sum = sum + v
	}
	fmt.Println("sum :", sum)

	for i, x := range x {
		if x == 3 {
			fmt.Println("index:", i)
		}
	}

	mobiles := map[string]string{"apple": "iPhone", "samsung": "galaxy", "google": "pixel"}
	for i, v := range mobiles {
		fmt.Printf("%s -> %s\n", i, v)
	}

	for v := range mobiles {
		fmt.Println("keys :", v)
	}

	for i, v := range "go" { //strings and runes
		fmt.Println(i, v)
	}

}
