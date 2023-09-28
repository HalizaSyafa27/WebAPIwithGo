package utils

import (
	"math/rand"
	"time"
)

func HandleRequest() {

}

// GenerateID generates a random ID
func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 10
	id := make([]byte, length)
	for i := 0; i < length; i++ {
		id[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(id)
}
