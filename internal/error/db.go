package error

import "fmt"

type KeyAlreadyExist struct {
	Msg   string
	Field string
}

func (e *KeyAlreadyExist) Error() string {
	return fmt.Sprintf("key \"%s\" is not unique", e.Field)
}
