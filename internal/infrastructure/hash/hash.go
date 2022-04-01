package hash

import "golang.org/x/crypto/bcrypt"

func Make(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(bytes), err
}

func Check(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
