// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	ID int `json:"id"`
// 	Name string `json:"name"`
// 	Email string `json:"email"`
// }

// // Fake db
// var users = []User{{ID: 1, Name: "saalim", Email: "muhammedsaalim005@gmail.com"}}

// func main() {
// 	r := gin.Default()

// 	// Api grouping
// 	api := r.Group("/api")
// 	{
// 	api.GET("/user", getUser)
// 	api.POST("/users", createuser)
// 	}
// 	// server start
// 		r.Run(":8000")
// }

// // GET
// func getUser(c *gin.Context)  {
// 	c.JSON(200, gin.H{
// 		"users":users,
// 	})
// }

// //POST
// func createuser(c *gin.Context)  {
// 	var newUser User

// 	if err := c.ShouldBindJSON(&newUser); err !=nil{
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error" : err.Error(),
// 		})
// 		return
// 	}
// 	newUser.ID = len(users) +1
// 	users = append(users, newUser)
// fmt.Println(users)

// 	c.JSON(http.StatusCreated, gin.H{
// 		"message":"UserCreatedSuccessfully",
// 		"user": newUser,
// 	})
// }

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
// fake db
var users = []User{{Name: "saalim", Age: 12}}

func main() {

	server := gin.Default()
	//HTTP request handler
	//•	Router
	//•	Middleware manager
//	•	Context creator
//	•	Error handler

// the process that behind the Default  i can also create my own(.new)
// engine := gin.New()
// engine.Use(Logger())
// engine.Use(Recovery())
// return engine

//     --Grouping the API--
	api := server.Group("/api")
	{
		api.GET("/user", getUser)
		api.POST("/users", setUser)
	}
	server.Run(":8000")
}
// GET
// *gin.Context is the bridge
// between the HTTP world and your Go code.
func getUser(c *gin.Context)  {
	c.JSON(200, gin.H{
	//	c.JSON is a Gin helper method that does 4 things:
	// 1.	Sets HTTP status code
	// 2.	Sets Content-Type: application/json
	// 3.	Converts Go values → JSON
	// 4.	Writes JSON to response body

// ------gin.H-------bg

"users":users,
	//map[string]interface{}{
	//"users": users,}
	})
}

// Post
func setUser(c *gin.Context)  {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":newUser,
		})
		return
	}

	users = append(users, newUser)
	c.JSON(http.StatusCreated, gin.H{
		"message": "UserCreatedSuccessfully",
		"user": newUser,
	})
}

//			       -- Explanation --
// The handler receives a *gin.Context which encapsulates 
// the HTTP request and response. Calling c.JSON sets
// the status code, serializes Go data into JSON, 
// sets headers, and writes the response back to the client.

