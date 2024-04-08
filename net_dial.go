package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	var a string
	fmt.Sprintf("hello", a)

	fmt.Println(a)

	r, err := net.Dial("tcp", "golang.org:80")
//net.Dial connects to server
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)

	r2, err := net.Listen("tcp", ":8080")
	//net.Listen creates server
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r2)

}
