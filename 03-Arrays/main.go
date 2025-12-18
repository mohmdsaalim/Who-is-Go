package main

import (
	"fmt"
)
// array is collection of same type of data type 
func main() {
// Inputing a value through terminal using fmt
fmt.Print("Enter the name of player ")
var input string
fmt.Scan(&input)//taking the Value form the terminal
	// this is basic array declaration in GO lang..
	footballers := [5]string{"Messi","Mbappe"}
	footballers[2] = "Neymar"
	footballers[3] = "Pedri"
	footballers[0] = input // directly change becouse of 
	fmt.Println(footballers[1:3])// sclicing from an array

// itraite the array using for loop
	for v,i := range footballers{
		if v%2 == 0{
				// continue
				fmt.Println(i,v)
		}
	}

}// cannot append
