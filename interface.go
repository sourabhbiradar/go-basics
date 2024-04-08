package main

import (
	"fmt"
	"math"
)

type geo interface {
	area() float32
	perimeter() float32
}
type rect struct {
	width, length float32
}

type circle struct {
	radius float32
}

func (r rect) area() float32 { // cant use (r *rect)
	return r.width * r.length
}
func (r rect) perimeter() float32 {
	return 2*r.width + 2*r.length
}
func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius

}
func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}
func ans(g geo) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
func main() {
	fmt.Println("interface")

	p := rect{width: 10, length: 5}
	c := circle{radius: 5}

	ans(p)
	/*fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())*/
	ans(c)
	/*fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())*/

}
