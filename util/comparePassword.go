package util

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}
