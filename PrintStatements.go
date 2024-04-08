package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("%T %v %q", 100, 100, "hero") //no new line ;Printf(%)
	fmt.Println("\ngodaddy" + "." + "com")   //Print'ln' ;new line
	fmt.Print("go", "by", "examples", ".", "com\n")

	var x = fmt.Sprintln(8) //Sprintln() assigns value to variable
	fmt.Println(x)
	r := fmt.Sprintln("help")
	fmt.Println(r)

	log.Println(": Date & time ") //yyyy/mm/dd  hr:min:sec

}
