package crypto

import (
	"crypto/rand"
	"fmt"
)

func GenerateRandomToken16() (string, error) {
	return GenerateRandomToken(16)
}

func GenerateRandomToken32() (string, error) {
	return GenerateRandomToken(32)
}

func GenerateRandomToken(n int) (string, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)
	// %x	base 16, lower-case, two characters per byte
	return fmt.Sprintf("%x", token), err

}
