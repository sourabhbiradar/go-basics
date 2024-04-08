package main

import (
	"fmt"
)

type t20 struct {
	season  int
	teams   []string
	sponser string
}

func main() {
	fmt.Println("advanced make map")

	var ipl t20

	var t []t20 = make([]t20, 3)
	t[0].teams = []string{"RCB", "KKR"}
	t[1].teams = []string{"CSK"}
	t[2].teams = []string{"MI"}
	fmt.Println(t)

	ipl.sponser = "TATA"
	ipl.season = 15
	ipl.teams = []string{"RCB", "MI", "CSK", "KKR"}

	fmt.Println(ipl)

	var winners map[string]interface{} = make(map[string]interface{})
	winners["MI"] = t20{13, []string{"Sachin", "Ponting", "Malinga"}, "TATA"}
	fmt.Println(winners)
}
