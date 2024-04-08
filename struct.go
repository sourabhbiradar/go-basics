package main

import (
	"fmt"
)

type student struct {
	name  string
	marks int
}

/*func (s *student) newStudent(name string){
}*/
func newStudent(names string) *student {
	s := student{name: names}
	//var s student
	//s.name = names
	s.marks = 90
	return &s

}

func main() {
	fmt.Println("structure")

	fmt.Println(student{"abhi", 45}) //breaking encaps
	fmt.Println(student{name: "Yash", marks: 33})
	fmt.Println(student{name: "krtk"})
	fmt.Println(&student{name: "aksh", marks: 67})
	fmt.Println(newStudent("abcd"))

	s := student{name: "wxyz", marks: 70} //abstraction
	fmt.Println(s)
	fmt.Println(s.name)
	fmt.Println(s.name, s.marks)

	sm := &s
	fmt.Println(sm.marks)

	sm.marks = 65
	fmt.Println(sm.marks) //struct are mutable

}
