package session

import (
	"math/rand"
	"strconv"
	"time"
)

var sessions = make(map[string]string)

func CreateSession(username string) string {
	rand.Seed(time.Now().UnixNano())
	sessionID := strconv.Itoa(rand.Int())
	sessions[sessionID] = username
	return sessionID
}

func GetUser(sessionID string) (string, bool) {
	user, ok := sessions[sessionID]
	return user, ok
}

func DeleteSession(sessionID string) {
	delete(sessions, sessionID)
}