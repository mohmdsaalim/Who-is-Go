package main

import "fmt"

func main() {
	a, b := 1,2
	v := &a// copied the adress of a //    point to ar 
	*v = 10 // changed the value of the a that we can change with addres location
	fmt.Println(*v,b)
}