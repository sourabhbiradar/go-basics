package main

import (
	"fmt"
)

func main() {
	fmt.Println("pointers to datatype")

	arr := [4]string{"a", "b", "c", "d"}
	arr1 := &arr
	var arr2 *[4]string = &arr
	fmt.Println(arr, arr1, arr2)

	s := []string{"a", "b", "c", "d"}
	s1 := &s
	var s2 *[]string = &s
	fmt.Println(s, s1, s2)

	m := map[string]int{
		"abc": 1,
	}
	m1 := &m
	var m2 *map[string]int = &m
	fmt.Println(m, m1, m2)
}
