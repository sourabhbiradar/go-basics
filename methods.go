package main

import (
	"fmt"
)

type rect struct {
	width, length int
}

func (r *rect) perimeter() int { //has reciever of type '*rect'
	return 2*r.width + 2*r.length
}
func (r rect) area() int { //defined for pointers and value recievers
	return r.width * r.length
}

func main() {
	fmt.Println("Methods")

	r := rect{width: 10, length: 5} //??
	fmt.Println("area :", r.area())
	fmt.Println("perimeter :", r.perimeter())

	r1 := &r
	fmt.Println(r1.area())
	fmt.Println(r1.perimeter())

}
