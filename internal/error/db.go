package error

import "fmt"

type KeyAlreadyExist struct {
	Msg   string
	Field string
}

func (e *KeyAlreadyExist) Error() string {
	return fmt.Sprintf("key \"%s\" is not unique", e.Field)
}

type KeyNotFound struct {
	Msg   string
	Field string
}

func (e *KeyNotFound) Error() string {
	return fmt.Sprintf("key \"%s\" is not found", e.Field)
}

type ForbiddenAction struct {
	Msg   string
	Field string
}

func (e *ForbiddenAction) Error() string {
	return fmt.Sprintf("forbidden action \"%s\" for \"%s\"", e.Msg, e.Field)
}
