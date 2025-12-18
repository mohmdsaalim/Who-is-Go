package main

import (
	"fmt"

)
func main() {
// slice of structs
// Slice have contain the backing arrays defultly it is 6 or capis 6
// 0. value of slice is nill
arr := []struct {
 i int
 o string
}{
{10, "Saalim"},
{20, "Messi"},
}

names := []int{1,2,3,4,5,6}
names = append(names, 3)

fmt.Println(names, arr)

var a []int
fmt.Printf("the len %d cap %d", len(a), cap(a))

//using make funtion 
// make is a  func that provided by go 
// to create refrence type like the slice and structs 
aa := make([]int, 0)// second argmend is length  
bb := make([]int , 0, 5)// third argment is capacity 
aa = append(aa, 1,12,12)
fmt.Println(aa)
fmt.Printf("%d", cap(bb))

	nestedValue := [][]string{
		[]string{"1","2","3"},
	    []string{"1","2","3"},
		[]string{"1","2","3"},
		[]string{"1","2","3"},
	}

	fmt.Println(nestedValue)
}
// Dynamic array 