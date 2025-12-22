package main

import "fmt"

type Students struct{
	Name string
	Age int
}


var studentslist []Students
func main() {
	for {
		fmt.Println("1. Add student")
		fmt.Println("2. Delete student")
		fmt.Println("3. List student")

		var input int 
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Print("Enter the name ")
			var name string
			fmt.Scan(&name)
			fmt.Print("Enter the age ")
			var age int
			fmt.Scan(&age)
			Addstudents(name, age)
		case 2:
			list()
			fmt.Println("Enter the index to delete")
			var index int
			fmt.Scan(&index)
			Delete(index)
		case 3:
			list()

		}
	}
}
// add
func Addstudents(name string, age int){
	studentslist = append(studentslist, Students{Name: name, Age: age})
}
//list
func list(){
	if len(studentslist)==0{
		fmt.Println("No students fount")
	}else{
		for i,v := range studentslist{
			fmt.Println(i, v.Name, v.Age)
		}
	}

}

// Delete
func Delete(index int){
studentslist = append(studentslist[:index], studentslist[index+1:]... )
fmt.Println("deleted")
}

