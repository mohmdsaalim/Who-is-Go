// interfaces are contracts 

package main

import "fmt"

// 1️⃣ Interface (contract)
type Payment interface {
	Pay(amount int)
}

// 2️⃣ CreditCard struct
type CreditCard struct {
	cardNumber string
}

// CreditCard implements Payment interface
func (c CreditCard) Pay(amount int) {
	fmt.Println("Paid", amount, "using Credit Card")
}

// 3️⃣ UPI struct
type UPI struct {
	upiID string
}

// UPI implements Payment interface
func (u UPI) Pay(amount int) {
	fmt.Println("Paid", amount, "using UPI")
}

// 4️⃣ Function that uses the interface
func makePayment(p Payment, amount int) {
	p.Pay(amount)
}

// 5️⃣ main function
func main() {

	cc := CreditCard{cardNumber: "1234-5678"}
	upi := UPI{upiID: "saalim@upi"}

	makePayment(cc, 500)
	makePayment(upi, 1000)
}