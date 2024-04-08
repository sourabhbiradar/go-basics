package main

import (
	"fmt"
)

func changeVal1(str string) {
	str = "hello"
	fmt.Println(str)
}
func changeVal2(str *string) {
	//*str = "helloCoder"
	fmt.Println(str)  //str==&tochange==location
	fmt.Println(*str) //*str==dereference &tochange==data

}

func main() {
	fmt.Println("Pointers and dereference operators")

	x := 7
	y := &x
	fmt.Println(x, y)

	*y = 8
	fmt.Println(x, y)

	toChange := "hey!!"
	fmt.Println(toChange)
	changeVal1(toChange)
	fmt.Println(&toChange) //location
	changeVal2(&toChange)  //location and if derefernecd('*str')==data

	var pointer *string = &toChange
	//var pointer *string=toChange // cant use toChange as type *string in variable declaration
	fmt.Println(pointer)  //address of toChange
	fmt.Println(*pointer) //deref &toChange so data
	fmt.Println(pointer, &pointer)
	fmt.Println(&pointer) //address of address of toChange

}
