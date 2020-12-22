package validator

import (
	"errors"
	"strings"
)

func ValidateRequired(field string, message string) error {
	if strings.TrimSpace(field) == "" {
		return errors.New(message)
	}
	return nil
}
