package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)




type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	// GET METHOD
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			user := User{Name: "Saalim", Age: 21}
			json.NewEncoder(w).Encode(user)
			return
		}
		//POST METHOD
		if r.Method == http.MethodPost {
			var user User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			// Business Logic 
			user.Age += 1

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
		http.Error(w, "Methode is not allowed", http.StatusMethodNotAllowed)
	})
	fmt.Println("Server is running")
	http.ListenAndServe(":8080", nil)

}