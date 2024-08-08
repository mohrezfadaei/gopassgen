package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	defaultPasswordLength = 16
	uppercaseLetters      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters      = "abcdefghijklmnopqrstuvwxyz"
	digits                = "0123456789"
	symbols               = "$@#$%^&*()_+{}|:\"<>?,./~`"
)

func GeneratePassword(length int, excludeUppercase, excludeLowercase, excludeDigits, excludeSymbols bool, includeChars string) (string, error) {
	if length <= 0 {
		length = defaultPasswordLength
	}

	if excludeUppercase && excludeLowercase && excludeDigits && excludeSymbols && includeChars == "" {
		return "", errors.New("no characters to generate password")
	}

	var charset string

	if !excludeUppercase {
		charset += uppercaseLetters
	}
	if !excludeLowercase {
		charset += lowercaseLetters
	}
	if !excludeDigits {
		charset += digits
	}
	if !excludeSymbols {
		charset += symbols
	}
	if includeChars != "" {
		charset += includeChars
	}
	charset += includeChars

	if charset == "" {
		return "", errors.New("no characters available to generate password")
	}

	var password strings.Builder
	for i := 0; i < length; i++ {
		char, err := randomChar(charset)
		if err != nil {
			return "", err
		}
		password.WriteByte(char)
	}

	return password.String(), nil
}

func randomChar(charset string) (byte, error) {
	max := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return charset[n.Int64()], nil
}
