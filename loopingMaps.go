package main

import (
	"fmt"
)

func printMap(c map[string]string) {
	for i, v := range c {
		fmt.Println("Hex code for", i, "is", v)
	}

}

func main() {
	fmt.Println("Looping Maps")
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)

}
