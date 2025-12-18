package structs

import "fmt"

func main(){

	User1 := Names{"Saalim", 19}
	a := &User1
	a.Name = "Mbappe"
	fmt.Println(*a)

}
// struct are native type of Go
// we can store anything on here it can be int, string, etc.
type Names struct{
	Name string
	age int

}