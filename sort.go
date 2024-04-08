package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("sorting")

	x := []int{4, 6, 2, 88, -1, 0}
	(sort.Ints(x))
	fmt.Println((x))

	y := []string{"cat", "dog", "rat", "cow", "fox"}
	sort.Strings(y)
	fmt.Println(y)

	z := []float64{9.8, 0.7, -0.77, -88.6, 104.2}
	sort.Float64s(z)    // only float64 can b sorted 'Float64s'
	fmt.Println(z)

	

}
