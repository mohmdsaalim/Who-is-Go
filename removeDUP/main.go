package main

import "fmt"

func main() {
	arr := []int{1,2,2,3,3,4,4,5,5,1}

	new := rem(arr)
	fmt.Println(new)

}

func rem(nw []int) []int {
	seen := make(map[int]bool)
	result :=[]int{}

	for _,val := range nw{
		if !seen[val]{
			seen[val] = true
			result = append(result, val)
		}
	}
return result
}