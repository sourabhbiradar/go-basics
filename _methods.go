package main

import (
	"fmt"
)

type Student struct {
	name  string
	marks []int
	age   int
}

func (s Student) getAge() int { //jst to fetch data (s Student)
	return s.age
}
func (s *Student) setAge(age int) int { //to modify data (s *Student)
	s.age = age
	//fmt.Println(age)
	return age
}
func (s *Student) getAvg() float64 {
	sum := 0

	for _, v := range s.marks {
		fmt.Println(v)
		sum = sum + v //sum=0+70->sum=70+90->sum=160+85->sum=245 once thrs nthing to itirate loop exits

	}
	return float64(sum) / float64(len(s.marks))

}
func (s *Student) getMaxMarks() int {
	curMarks := 0
	for _, v := range s.marks {
		if curMarks < v {
			curMarks = v
		}
	}
	return curMarks
}
func main() {
	fmt.Println("Methods")
	s1 := Student{"Abc", []int{70, 90, 85}, 20}
	s2 := Student{"Xyz", []int{73, 96, 81}, 24}
	//s1.getAge()
	fmt.Println(s1.getAge())

	//	fmt.Println(s1.setAge(33), s1)
	s2.setAge(29)
	fmt.Println(s2)

	fmt.Println(s1.getAvg())
	avg2 := s2.getAvg()
	fmt.Println(avg2)

	maxMarks := s1.getMaxMarks()
	fmt.Println(maxMarks)
	fmt.Println(s2.getMaxMarks())

}
