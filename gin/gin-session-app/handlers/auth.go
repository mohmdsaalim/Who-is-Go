package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gin-session-app/session"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Hardcoded credentials
	if req.Username != "saalim" || req.Password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	sessionID := session.CreateSession(req.Username)

	// Set cookie
	c.SetCookie("session_id", sessionID, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func Logout(c *gin.Context) {
	sessionID, _ := c.Cookie("session_id")
	session.DeleteSession(sessionID)

	c.SetCookie("session_id", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func Profile(c *gin.Context) {
	user := c.GetString("user")
	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"msg":  "welcome to protected route",
	})
}