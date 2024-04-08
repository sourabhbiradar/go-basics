package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("string conversion")
	x := "hi hey hello"
	fmt.Println([]byte(x)) //letter by letter to byte acc. to ascii table
	fmt.Println([]rune(x))

	y := []string{"hi", "hey", "hello"}
	fmt.Println(strings.Join((y), ",")) //cov []string to string n join with ','

    

}
