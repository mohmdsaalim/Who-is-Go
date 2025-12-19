
package main

import "fmt"

type user struct{
	name , age string
}
// MAP with make syntax
func main(){
	m := make(map[string]user)
	m["name"] = user{"saalim", "12"}

	
	// Litral Syntax
	b := map[string]user{
		"names" : user{"messi", "12"},
		"bub" : user{"cristiano", "21"},
	}
	fmt.Println(m,b)
}