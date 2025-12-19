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

names := []int{1,2,3,4,5,6}// slice Litral syntax
names = append(names, 3)

fmt.Println(names, arr)

var a []int
fmt.Printf("the len %d cap %d", len(a), cap(a))

//using make funtion 
// make is a  func that provided by go 
// to create refrence type like the slice and structs 
aa := make([]int, 0)// second argmend is length  // MAKE FUNC SYNTAX
bb := make([]int , 0, 5)// third argment is capacity 
aa = append(aa, 1,12,12)
fmt.Println(aa)
fmt.Printf("%d", cap(bb))

// THE CAPACITY IS DOUBLE WHEN WE ADD SINGLE ELEMENTS
slice := [][]int{
	[]int{1,2,3,4,5},
	[]int{2,3,4,5,6},
}

fmt.Println(slice)
	nestedValue := [][]string{
		[]string{"1","2","3"},
	    []string{"1","2","3"},
		[]string{"1","2","3"},
		[]string{"1","2","3"},
	}

	fmt.Println(nestedValue)




  slices := make([]int, 0,5)
	for v,i := range slices{
		slices = append(slices, v, i)
		fmt.Println(slices)
	}
}
// Slice is  Dynamic array we can extend, remove, add, slice tec..
//length 1, capacity 4
//length 2, capacity 4
//length 3, capacity 4
//length 4, capacity 4
//length 5, capacity 8
//length 6, capacity 8
//length 7, capacity 8
//length 8, capacity 8
//length 9, capacity 16
//length 10, capacity 16
//length 11, capacity 16
//length 12, capacity 16
//length 13, capacity 16
//length 14, capacity 16
//length 15, capacity 16
//length 16, capacity 16
//length 17, capacity 32
//length 18, capacity 32
//length 19, capacity 32
//length 20, capacity 32