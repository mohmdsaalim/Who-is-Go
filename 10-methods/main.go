package main

import "fmt"
//A method in Go is a function that is associated with a specific type, defined using a receiver.
type User struct{
	name string
	age int
}
// methods 
// we can assign a function to types like struct
//recever ⬇️type
// func (u User) call() {
// 	fmt.Println("helo welcome",u.name)
// }

func (a *User) super(){ //pppp
	a.name = "Saalim"//pppp
}

func main() {
	// fmt.Printf("Enter the name")
	// var input string
	// fmt.Scan(&input)
	// us := User{input, 23}
	// us.call()

newuser := User{"Messi", 12}//pppp
newuser.super()//pppp
fmt.Println(newuser)//ppp

}


// to reflect the org us ethe pointer recever
