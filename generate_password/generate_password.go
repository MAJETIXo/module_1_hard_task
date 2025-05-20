package generate_password

import (
	"crypto/rand"
	"math/big"
)

// letters - список допустимых символов в пароле
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GeneratePassword(n int) (string, error) {
	b := make([]byte, n)
	for i := range b {
		passw, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		b[i] = letters[passw.Int64()]
	}
	return string(b), nil
}
