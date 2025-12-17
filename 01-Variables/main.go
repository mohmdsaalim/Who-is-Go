// Variable declaration in Go Lang
package main

import "fmt"
// when i put caps letter in variable name it is exported 
// which means we can acsses that value to another files(import)
var A int = 100 // also we can put string, bool, float. etc 

func main() {
	var a int = 10 // normal way
	// simlpe methode to declare one or more var in one line 
	var b, c, d int = 1,2,3 //only same Datatype
	// declaring a variable but not assign a value
	var future int // outwill default value of that Data type
	future = 100 // we can assgn like this
	fmt.Println(a,b,c,d,future)

// declaring a varible with walrus operator //MOST COMMON WAY IN GO
	name := "saalim" //doesn't declare any var or datatype
	fmt.Println(name) 

// CONSTANTS IN GO
	const footballers string = "Messi"
//Go supports constants of character, string, boolean, and numeric values.

}