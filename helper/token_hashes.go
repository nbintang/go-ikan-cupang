package helper

import "golang.org/x/crypto/bcrypt"

func HashToken(token string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	return string(hash)
}


func CompareToken(token string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(token)) == nil
}