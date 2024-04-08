package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("buffered chan")
	c := make(chan int, 3)
	c <- 7
	c <- 8
	c <- 9
	go func(c chan int) {
		fmt.Println(<-c)
		fmt.Println(<-c)
	}(c)
	go func(c chan int) {
		fmt.Println(<-c)
	}(c)
	time.Sleep(3)

}
