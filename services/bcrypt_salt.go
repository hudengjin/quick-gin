package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// BcryptPassword bcrypt 算法 hash加 salt 加密密码
func BcryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// IsCorrectPassword 判断未加密的密码和 hash 加盐密码是否一致
func IsCorrectPassword(hashedPwd, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}