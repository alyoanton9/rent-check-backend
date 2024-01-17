package encrypt

import "golang.org/x/crypto/bcrypt"

type BCryptHash struct {
}

func (BCryptHash) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (BCryptHash) PasswordsMatch(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
