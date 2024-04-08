package main

import (
	"fmt"
)
type info struct{
	phone int
}
type login struct{
	phone
	userid int
	password string
}

func (p *info)number(){
	var ph int
	p.phone=ph
	fmt.Println("Enter phone number")
	fmt.Scanln(&ph)

}
func (l *login)signin(){
	var sp string
	l.password=ps
	l.userid=p.phone
	l.userid=ur

	if ur==ph {
		fmt.Println("Create Password")
		fmt.Scanln(&ps)
		fmt.Println("Confirm password")
		fmt.Scanln(&sp)
		if ps==sp{
			fmt.Println("Password created successfully")
		}else {
			fmt.Println("Enter characters don't match")
			fmt.Println("Try again!!")
		}
		}

}

func main() {
	fmt.Println("Login Page")
}
