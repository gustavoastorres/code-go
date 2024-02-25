package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Tuesday?")
	today := time.Now().Weekday()
	switch time.Tuesday {
	case today:
		fmt.Println("Today.")
	case (today + 1) % 7:
		fmt.Println("Tomorrow.")
	case (today + 2) % 7:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
