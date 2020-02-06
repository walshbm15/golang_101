package main

import (
	"fmt"
)

type player struct {
	name	string
	age		int
}

func main() {
	playerSlice := make([]player, 1)  // sets length to 1
	p1 := player{"Alex", 20}
	p2 := player{"Mark", 29}

	fmt.Println(len(playerSlice))  // len(playerSlice)=1
	playerSlice = append(playerSlice, p1, p2)  // Appends p1 and p2 to playerSlice
	
	fmt.Println(len(playerSlice))  // len(playerSlice)=3
	fmt.Println(playerSlice)

	for i, p := range playerSlice {  // iterate over range of slice
		fmt.Printf("player%d: %v\n", i, p)
	}


	playerMap := make(map[string]player)  // Create player map
	playerMap["player1"] = p1  // Add to map
	playerMap["player2"] = p2
	fmt.Println(playerMap)
}
