// it is a light weight thraed 
// managed by scheduler
package main
//CONCON
// conconrency managing execution of diffrnt task efficitantly sp doesnt block the system flow
// at a time one task and no order
import (
	"fmt"
	"time"
)
var A string = "kundan"
func main() {
	go name()
	 sathar()
	 inshad()

	time.Sleep(time.Second)
	fmt.Println("hi i am main function")
}

func name(){
	fmt.Println("hi i am go routine")

}