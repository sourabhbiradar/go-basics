package main

import (
	"fmt"
)

func val(i int) {

}
func pntr(iptr *int) {
	*iptr = 89
}

func main() {
	fmt.Println("Pointer")

	i := 3
	val(i)
	fmt.Println(i)
	pntr(&i)
	fmt.Println(i)
	fmt.Println(&i)

}
