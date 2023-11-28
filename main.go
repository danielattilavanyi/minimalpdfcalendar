package main

import (
	"fmt"

	"github.com/danielattilavanyi/minimalpdfcalendar/setup"
)

func main() {
	for i := 0; i < 12; i++ {
		fmt.Println(setup.Months[i])
	}
	fmt.Print("Horizontal Resolution: ")
	fmt.Println(setup.Horizontal_Resolution)
	fmt.Print("Vertical Resolution: ")
	fmt.Println(setup.Vertical_Resolution)
}
