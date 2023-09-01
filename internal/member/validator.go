package member

import (
	"fmt"
)

const NameMinLength = 3
const NameMaxLength = 50

var (
	ErrNameInvalidLength = fmt.Errorf(
		"member name must between %d-%d characters",
		NameMinLength,
		NameMaxLength,
	)
)

func ValidateName(name string) error {
	l := len(name)
	if l < NameMinLength || l > NameMaxLength {
		return ErrNameInvalidLength
	}
	return nil
}
