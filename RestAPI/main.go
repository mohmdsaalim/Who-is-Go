package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"Name":"messi",
			"age":12,
		})
		
	})
	
	fmt.Println("server RUN")
	http.ListenAndServe(":8000", nil)
}
// "encoding/json"
// "fmt"
// "net/http"
// server()
	// http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
	// 	json.NewEncoder(w).Encode(map[string]any{
	// 		"Name":"saalim",
	// 		"age":12,
	// 		"adress":"palakkad",
	// 	})
	// })
	
	// http.ListenAndServe(":5000", nil)

// func server()  {
// 	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
// 		json.NewEncoder(w).Encode(map[string]any{
// 			"Name":"saalim",
// 			"age":12,
// 			"adress":"palakkad",
// 		})
// 	})

// 	http.ListenAndServe(":5000", nil)
// }


