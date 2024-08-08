package generator

import (
	"crypto/rand"
	"math/big"
)

const (
	passwordLength = 16
	charset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GeneratePassword() (string, error) {
	var password []byte
	for i := 0; i < passwordLength; i++ {
		char, err := randomChar(charset)
		if err != nil {
			return "", err
		}
		password = append(password, char)
	}
	return string(password), nil
}

func randomChar(charset string) (byte, error) {
	max := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return charset[n.Int64()], nil
}
