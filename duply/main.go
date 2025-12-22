package main

import "fmt"

func main() {
arr := [6]int{1000,2,3,4,5,600}
max := arr[0]
min := arr[0]

for i := 0; i < len(arr); i++{
	if max < arr[i]{
		max = arr[i]
	}
	if min > arr[i]{
		min = arr[i]
	}
}
fmt.Println("max is ",max)
fmt.Println("min is ",min)
}