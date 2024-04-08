package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func constants() {
	const q = 2
	//q=5
	fmt.Println(q)

	const r string = "artist"
	fmt.Println(r)

	const a = 1234.
	const f = a / 4.
	fmt.Println(f)
	
	const p = 11e3 //11 x 10 pow2=>11x10x10x10
	fmt.Println(p)
	fmt.Println(math.Pi)
	fmt.Println(math.Tan(45))
	fmt.Println(math.Cos(45))
	fmt.Println(math.Sin(45))
	fmt.Println(math.Sin(45) / math.Tan(45))
	fmt.Println(cmplx.Cot(45))
}
func main() {
	constants()
}
