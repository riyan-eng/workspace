package util

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

func GenerateHash(str *string) (string, error) {
	hashedStr, err := bcrypt.GenerateFromPassword([]byte(*str), bcrypt.DefaultCost)
	return string(hashedStr), err
}

func VerifyHash(hashedStr, candidateStr string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(candidateStr)); err == nil {
		return true
	} else {
		return false
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
