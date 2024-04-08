package main

import (
	"fmt"
)

type Point struct {
	x, y int32
}
type Circle struct {
	radius float64
	origin *Point
	//origin Point
	//*Point //for &Point{5,2}
	//Point //for Point{5,2}
}

func main() {
	fmt.Println("YT inheritence")
	c1 := Circle{3.67, &Point{5, 2}}
	//c1 := Circle{3.67, Point{5, 2}}
	fmt.Println(c1)
	fmt.Println(c1.radius)
	fmt.Println(c1.origin)
	fmt.Println(c1.origin.x)
	fmt.Println(c1.origin.y)
	//fmt.Println(c1.x)
	//fmt.Println(c1.y)
}
