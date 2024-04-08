package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map")
	m := map[int]string{
		12: "apples",
		5:  "bananas",
		4:  "mangos", //',' comma at last value impt.
	}
	fmt.Println(m[5])
	m[8] = "figs" //no need to use append for adding new element
	fmt.Println(m)

	fmt.Println(len(m))

	delete(m, 5) // delete(map_name,key)
	fmt.Println(m)

	//value, ok := m[12]//prints value of m[key] if exists with ok==true

	value, ok := m[99] //prints flase as no key/value in m[99]
	fmt.Println(value, ok)

	fmt.Println(len(m))

}
