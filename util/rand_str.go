package util

import "math/rand"

// Charset is the default charset for RandomString.
const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a random string of length `length`.
func RandomString(length uint) string {
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = Charset[rand.Intn(len(Charset))]
	}
	return string(runes)
}
