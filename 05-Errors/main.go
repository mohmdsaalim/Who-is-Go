package main

import "fmt"

func main() {
	result, err := divide(10,2)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("The value is ",result)
	}
}

func divide(a, b int) (int, error){
if b == 0 {
	return 0, fmt.Errorf("The value cant divided by %d",b)
}
return a/b, nil
}