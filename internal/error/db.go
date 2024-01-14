package error

import "fmt"

type KeyAlreadyExist struct {
	Msg   string `default:"unique"`
	Field string
}

func (e *KeyAlreadyExist) Error() string {
	return fmt.Sprintf("key \"%s\" is not unique", e.Field)
}

type KeyNotFound struct {
	Msg   string `default:"not-found"`
	Field string
}

func (e *KeyNotFound) Error() string {
	return fmt.Sprintf("key \"%s\" is not found", e.Field)
}

type NoAccess struct {
	Msg   string `default:"access"`
	Field string
}

func (e *NoAccess) Error() string {
	return fmt.Sprintf("no access: \"%s\"", e.Field)
}
