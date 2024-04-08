package main

import (
	"fmt"
	"time"
)

func chansync(m chan int) {
	fmt.Println("Working")
	//m <- 989
	time.Sleep(time.Second)
	//m <- 989
	fmt.Println("Done")
	//fmt.Println(<-m)

	m <- 989

}
func main() {
	fmt.Println("Channel Synchronization")
	m := make(chan int, 1)
	//	m <- 77
	go chansync(m)
	<-m
}
