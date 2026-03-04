package utils

import "golang.org/x/crypto/bcrypt"

// 生成hash密码
func GeneratePasswordHash(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(fromPassword)

	return hash, nil
}

// 比较密码
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
