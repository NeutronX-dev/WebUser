package WebUser

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func MakeToken(username, password string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	RandomStr := RandomString(rand.Intn(len(password)) + len(username))
	var res []byte = make([]byte, 0)
	for i := 0; i < len(username); i++ {
		if len(password)-1 >= i {
			res = append(res, password[i])
		}
		if len(RandomStr)-1 >= i {
			res = append(res, RandomStr[i])
		}
		res = append(res, username[i])
	}
	return string(res)
}

func Hash256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return string(hex.EncodeToString(h.Sum(nil)))
}
