package logic

import (
	"time"
	"math/rand"
)

// init initialises rand's Seed when the logic package is called 
func init() {
    rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// RandStr generates a alphanumeric string of size "size"
func RandStr(size int) string {
    b := make([]rune, size)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}