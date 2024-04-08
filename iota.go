package main

import (
	"fmt"
)

func main() {
	// iota is ninth letter in greek
	// sequence of numbers default frm zero
	// only be used with const

	fmt.Println("iota")

	const x = iota
	fmt.Println(x)

	const (
		C0 = iota
		C1 = iota
		C2 = iota
	)
	fmt.Println(C0, C1, C2)

	const (
		i0 = iota - 1 //0 - 1 == -1 ==i0
		i1
		_ // skips i2 == 1
		i3
		i4
		i5
	)
	fmt.Println(i0, i1, i3, i4, i5) // -1 0 2 3 4

	const (
		a = iota + 12 //0 + 12 == 12 == x0
		b
		c
		d
	)
	fmt.Println(a, b, c, d) // 12 13 14 15

}
