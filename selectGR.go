package main

import (
	"fmt"
	"time"
)

func port1(chan1 chan string) {
	time.Sleep(1)
	chan1 <- "Port1 entered"
}
func port2(chan2 chan string) {
	time.Sleep(1)
	chan2 <- "Port2 entered"
}
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go port1(c1)
	go port2(c2)

	select {
	case op1 := <-c1:
		fmt.Println(op1)

	case op2 := <-c2:
		fmt.Println(op2)
	}
}
