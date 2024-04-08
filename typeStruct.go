package main

import (
	"fmt"
)

type Point struct { //first letter CAPS (naming convention)
	x int32
	y float32
	z int32
}

func changeX(pt *Point) {
	pt.y = 11.43
	fmt.Println(pt) //==p3
}

func main() {
	fmt.Println("struct/custom type")

	p1 := &Point{3, 4.2, 7}
	p2 := Point{4, 6.5, 2}
	//p2 := Point{4, 6.5}//too few values in struct literal
	fmt.Println(p1)
	fmt.Println(p2)

	p3 := &Point{x: 3, z: 5}
	fmt.Println(p3)
	fmt.Println(&p3)

	changeX(p3)
	fmt.Println(p3)

}
