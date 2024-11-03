package util

import (
	"math/rand"
	"time"
)

// RandomStringWithSample is a function
func RandomStringWithSample(n int, sample string) string {
	var letters = []rune(sample)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}
