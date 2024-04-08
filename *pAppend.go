package main

import "fmt"

func main() {
	fmt.Println("append *pointer")

	x := []string{"hello"}
	fmt.Println(x)
	x = append(x, "world")
	fmt.Println(x)

	var y *[]string = &x
	fmt.Println(*y)
	*y = append(*y, "!!")
	fmt.Println(*y, y, &y, &*y)

}
