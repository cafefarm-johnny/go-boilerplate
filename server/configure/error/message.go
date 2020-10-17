package error

import "fmt"

type errorMessage struct {
	Code    int
	Message string
}

func newErrorMessage(code int, message string) *errorMessage {
	em := new(errorMessage)
	em.Code = code
	em.Message = message

	return em
}

func (em *errorMessage) String() string {
	return fmt.Sprintf("code: %d, message: %s", em.Code, em.Message)
}
