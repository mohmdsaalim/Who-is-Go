package main

import (
	"fmt"
)

func main() {
// Inputing a value through terminal using fmt
fmt.Print("Enter the name of player ")
var input string
fmt.Scan(&input)//taking the Value form the terminal
	// this is basic array declaration in GO lang..
	footballers := [5]string{"Messi","Mbappe"}
	footballers[2] = "Neymar"
	footballers[3] = "Pedri"
	footballers[4] = input
	fmt.Println(footballers)

// itraite the array using for loop
	for v,i := range footballers{
		if v%2 == 0{
				// continue
				fmt.Println(i,v)
		}
	}
}
