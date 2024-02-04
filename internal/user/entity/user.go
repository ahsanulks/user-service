package entity

import (
	"fmt"
	"regexp"
	"time"

	customerror "github.com/SawitProRecruitment/UserService/internal/customError"
)

const (
	phoneNumberMinLen = 10
	phoneNumberMaxLen = 13
)

type User struct {
	ID          string
	FullName    string
	PhoneNumber string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUser(fullName, phoneNumber, password string) (*User, error) {
	user := &User{
		FullName:    fullName,
		PhoneNumber: phoneNumber,
		Password:    password,
	}

	if err := user.ValidatePhoneNumber(); err != nil {
		return nil, err
	}
	return user, nil
}

func (user User) ValidatePhoneNumber() error {
	validationError := customerror.NewValidationError(map[string]string{})
	if len(user.PhoneNumber) < phoneNumberMinLen || len(user.PhoneNumber) > phoneNumberMaxLen {
		validationError.AddError("phoneNumber", fmt.Sprintf("must be between %d and %d characters in length", phoneNumberMinLen, phoneNumberMaxLen))
	}

	phoneNumberRegex := `^\+62[0-9]$`
	match, _ := regexp.MatchString(phoneNumberRegex, user.PhoneNumber)
	if !match {
		validationError.AddError("phoneNumber", "must start with '+62' and only containt number")
	}

	if validationError.HasError() {
		return validationError
	}

	return nil
}
