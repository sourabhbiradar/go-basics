package main

import "fmt"

func For() {
	i := 1
	for i <= 4 {
		fmt.Println(i)
		i++
	}
	for j := 3; j <= 7; j++ { //intializing;condition;j+1(after)
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}
	mobiles := [5]string{"moto", "sony", "bb", "iphone"}
	for i, v := range mobiles {
		fmt.Println(i, v)
	}
	for g := 1; g < 10; g++ {
		if g*5 == 0 {
			continue //continue to next iteration of the loop
		}
		fmt.Println(g)

	}
}
func main() {
	For()
}
