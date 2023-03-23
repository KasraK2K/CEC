package helper

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randSource.Intn(len(charset))]
	}
	return string(b)
}
