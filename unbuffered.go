package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("unbuffered chan") 
	// as soon as msg is sent to channel there shud b asnother to recieve
	// unbuffered channel used to PURELY communicate with go rountines (i guess) 
	c := make(chan string)

	go func() {
		c <- "message recieved"

	}()
	go func(c chan string) {

		fmt.Println(<-c)
	}(c)
	time.Sleep(3)

}
