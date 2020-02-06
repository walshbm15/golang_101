package main

import (
	"fmt"
	"time"
)

func forLoop(){
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("forLoop() sum: %d\n", sum)
}

func whileLoop(){
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Printf("whileLoop() sum: %d\n", sum)

	if v := 5; v < 6 {
		fmt.Println("Statement is true")
	} else {
		fmt.Println("Statement is false")
	}

}

func saturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func main() {
	defer fmt.Println("Main has ended")

	forLoop()
	whileLoop()
	saturday()
}
