package main

import (
	"fmt"
	"time"
)

func main() {
	welcome := "Welcome to Time Handling"
	fmt.Println(welcome)

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01/02/2006, Monday, 15:04:05, MST"))
}
