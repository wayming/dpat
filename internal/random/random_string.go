package random

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandString(strSize int) string {
	bytes := make([]rune, strSize)
	for index := range bytes {
		bytes[index] = letters[rand.Intn(len(letters))]
	}

	return string(bytes)
}
