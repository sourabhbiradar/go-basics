package main

import (
	"fmt"
	"os"
)

func xit() {
	defer fmt.Println("hey") //wont execute becoz of os.Exit
	fmt.Println("hello")
	os.Exit(2) //when executed exit status prints by default
}

func main() {
	fmt.Println("EXIT")
	xit()

}
