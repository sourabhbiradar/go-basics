package main

import (
	"fmt"
)

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	fmt.Println("function closure")

	returnFunc("hello")()
	returnFunc("hey")() //==func()

}
