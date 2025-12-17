package main

import "fmt"

func main() {
	//There are different type of for Loops is here
	i := 1
	for i<=3{
		fmt.Println(i)
		i++
	}

	// most comman 
	for i := 1; i<=10; i++{
		fmt.Println("ten times", i)
	}

	// CONTINUE  in loops
	for n := range 6{
			if n%2==0{
				// fmt.Println(n)
				continue
			}
			fmt.Println(n)
			
	}
	// Range Loop
	// looping the map 
	footballers := make(map[string]any) //This is an map key value pairs
	footballers["JS"] = "Javascript"
		for v,i := range footballers{
			fmt.Println(i,v)
		}
}