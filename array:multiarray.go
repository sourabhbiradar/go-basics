package main

import "fmt"

func array() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	a[3] = 100
	fmt.Println(a)
	fmt.Printf(" Length of a[] =%v\n Capacity of a[]=%v\n", len(a), cap(a))
}
func multiarray() {
	y3D := [3][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 0, 1}}
	fmt.Println(y3D)
	fmt.Println(y3D[0][2]) //print frm '[0]==1st array'the '[2]==3rd element'

}

func main() {
	fmt.Println("ARRAY")
	array()
	fmt.Println("Multi Array")
	fmt.Println("Arrays are single dimension but we can create multiple dimension array!!")
	multiarray()

}
