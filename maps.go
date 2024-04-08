package main

import (
	"fmt"
)

func main() {
	fmt.Println("map")
	m := make(map[int]string)
	m[7] = "Ronaldo"
	m[18] = "Kohli"
	m[10] = "Modric"
	fmt.Println(m)

	x := m[7] // key is important
	fmt.Println("x =", x)

	fmt.Println("length :", len(m))

	delete(m, 10) //delete key/value using key
	fmt.Println(m)

	_, get := m[7] //true ??
	fmt.Println(get)

	_, fetch := m[10] //false
	fmt.Println(fetch)

	y := map[int]string{1: "KL Rahul", 2: "Bumrah"}
	fmt.Println(y)

	z := map[string]int{
		"CR": 7,
		"VK": 18,
		"LM": 10,
	}
	fmt.Println(z)
}
