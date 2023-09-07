package service

import "golang.org/x/crypto/bcrypt"

// HashPassword เข้ารหัสรหัสผ่านโดยใช้ bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword ตรวจสอบว่ารหัสผ่านตรงกับรหัสผ่านที่เข้ารหัส
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
