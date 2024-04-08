package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {
	fmt.Println("Go Routine")
	f("xyz")
	go f("abc")
	go func(m int) {
		fmt.Println(m)
	}(123)
	time.Sleep(time.Second)
	fmt.Println("Done")

}
