package encrypt

type Hasher interface {
	HashPassword(password string) (string, error)
	PasswordsMatch(password, hash string) bool
}

func NewHasher() Hasher {
	return BCryptHash{}
}
