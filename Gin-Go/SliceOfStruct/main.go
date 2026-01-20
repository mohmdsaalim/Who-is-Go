package sliceofstruct
// package main

import "fmt"

// Step 1: Define a struct
type User struct {
	ID   int
	Name string
	Age  int
}

func main() {

	// Step 2: Create an empty slice of structs
	users := []User{}

	// Step 3: Add data (append structs into slice)
	users = append(users, User{ID: 1, Name: "Ali", Age: 22})
	users = append(users, User{ID: 2, Name: "Sara", Age: 25})
	users = append(users, User{ID: 3, Name: "John", Age: 30})

	// Step 4: Print the full slice
	fmt.Println("All Users:", users)

	// Step 5: Access single struct from slice
	fmt.Println("First User Name:", users[0].Name)

	// Step 6: Update a value inside slice
	users[1].Age = 26

	// Step 7: Loop through slice of structs
	fmt.Println("\nUpdated Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}
