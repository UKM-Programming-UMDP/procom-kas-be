package utils

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyz"
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}
