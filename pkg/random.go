package cry

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomStringWithCharset( n int, charset string) string {
	b := make([]byte, n)
	cn := len(charset)
	for i := range b {
		b[i] = charset[seededRand.Intn(cn)]
	}
	return string(b)
}

func RandomString(n int) string {
	return RandomStringWithCharset(n, charset)

}