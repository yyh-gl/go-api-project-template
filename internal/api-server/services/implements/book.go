package interfaces

import (
	"errors"
	"strings"
)

type bookService struct {}

func (bookService) Index(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")
