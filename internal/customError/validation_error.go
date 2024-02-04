package customerror

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	errors map[string]string
}

func NewValidationError(errors map[string]string) *ValidationError {
	return &ValidationError{errors: errors}
}

func (ve *ValidationError) AddError(key, message string) {
	if ve.errors == nil {
		ve.errors = make(map[string]string)
	}
	if val, ok := ve.errors[key]; ok {
		ve.errors[key] = val + ";" + message
	} else {
		ve.errors[key] = message
	}
}

func (ve *ValidationError) Error() string {
	var errorMessages []string
	for key, msg := range ve.errors {
		errorMessages = append(errorMessages, fmt.Sprintf("%s: %s", key, msg))
	}
	return strings.Join(errorMessages, ";")
}

func (ve ValidationError) HasError() bool {
	return len(ve.errors) > 0
}
