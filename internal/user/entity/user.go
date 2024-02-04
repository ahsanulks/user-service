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
	fullNameMinLen    = 3
	fullNameMaxLen    = 60
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

	validationError := customerror.NewValidationError()
	if err := user.validatePhoneNumber(); err != nil {
		validationError.Merge(err)
	}

	if err := user.validateFullName(); err != nil {
		validationError.Merge(err)
	}

	if validationError.HasError() {
		return nil, validationError
	}

	return user, nil
}

func (user User) validatePhoneNumber() error {
	validationError := customerror.NewValidationError()
	if len(user.PhoneNumber) < phoneNumberMinLen || len(user.PhoneNumber) > phoneNumberMaxLen {
		validationError.AddError("phoneNumber", fmt.Sprintf("must be between %d and %d characters in length", phoneNumberMinLen, phoneNumberMaxLen))
	}

	phoneNumberRegex := `^\+62[0-9]`
	match, _ := regexp.MatchString(phoneNumberRegex, user.PhoneNumber)
	if !match {
		validationError.AddError("phoneNumber", "must start with '+62' and only containt number")
	}

	if validationError.HasError() {
		return validationError
	}

	return nil
}

func (user User) validateFullName() error {
	validationError := customerror.NewValidationError()
	if len(user.FullName) < fullNameMinLen || len(user.FullName) > fullNameMaxLen {
		validationError.AddError("fullName", fmt.Sprintf("must be between %d and %d characters in length", fullNameMinLen, fullNameMaxLen))
	}

	if validationError.HasError() {
		return validationError
	}

	return nil
}
