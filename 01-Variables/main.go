// Variable declaration in Go Lang
package main

import "fmt"
// three patterns
//var a int = 10 //1
//a := 10 //2
// var a = 10 // 3


//  IN GO THERE IS NO UNDIFINED STAGE WHEN WE DECLARE A VAIABLE IT TAKE THE DEFAULT VALUE IF THE TYPE
// when i put caps letter in variable name it is exported 
// which means we can acsses that value to another files(import)
var A int = 100 // also we can put string, bool, float. etc ,  only way to declare a variable outside a function. 
var B = "string"//also support this
func main() {
	// Go supporting UTF-8
	var a int = 10 // normal way
	// simlpe methode to declare one or more var in one line 
	var b, c, d int = 1,2,3 //only same Datatype
	// declaring a variable but not assign a value
	var future int // outwill default value of that Data type
	future = 100 // we can assgn like this
	fmt.Println(a,b,c,d,future)

// declaring a varible with walrus operator //MOST COMMON WAY IN GO
	name := "saalim" //doesn't declare any var or datatype .... useCase is when i know what is the initail value i can use walrus
	fmt.Println(name) 

// CONSTANTS IN GO
	const footballers string = "Messi" // constants are not chnagable 
	const messi = "messi"
//Go supports constants of character, string, boolean, and numeric values.

// TYPE CASTING
var bro int8 = 10
var bo int32 = int32(bro)
fmt.Println(bo)
}
// need to explicitly convert the type of the variable  