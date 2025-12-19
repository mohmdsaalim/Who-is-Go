package main

import "fmt"
// In go there must be a main func 
// any support anoimous func inside the main function

func main() {
	// DEFER
	fmt.Println("The Loop started")
	defer fmt.Println("finished")// DEFER is keyword that work like when the main func is finish execution then the defer will work
	// not only print most common use is closing the file when panics
	fmt.Println("the sum is",sum(10,20))


	for a := 0; a <= 10; a++{
		defer fmt.Println(a)
	}
	sum := func (aa, bb int) int {
		return aa+bb
	}
	fmt.Println(sum(10,10))
}

func sum(a, b int) int{
	return a+b


}