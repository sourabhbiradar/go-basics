package main

import (
	"fmt"
)

func main() {
	fmt.Println("Buffered Channels")//doesnt need receiver for buffered channel
	message := make(chan string, 4) //4 values needed
	message <- "hello"
	message <- "hey"
	message <- "hi"
	message <- "mmmm"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
