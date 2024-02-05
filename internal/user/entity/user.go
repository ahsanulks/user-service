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
	passwordMinLen    = 6
	passwordMaxLen    = 64
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

	if err := user.validatePassword(); err != nil {
		validationError.Merge(err)
	}

	if validationError.HasError() {
		return nil, validationError
	}

	return user, nil
}

func (user *User) UpdateProfile(fullName, phoneNumber *string) error {
	validationError := customerror.NewValidationError()
	if fullName == nil && phoneNumber == nil {
		validationError.AddError("fullName", "at least 1 field must be present")
		validationError.AddError("phoneNumber", "at least 1 field must be present")
		return validationError
	}

	if fullName != nil {
		user.FullName = *fullName
		if err := user.validateFullName(); err != nil {
			validationError.Merge(err)
		}
	}

	if phoneNumber != nil {
		user.PhoneNumber = *phoneNumber
		if err := user.validatePhoneNumber(); err != nil {
			validationError.Merge(err)
		}
	}

	if validationError.HasError() {
		return validationError
	}

	return nil
}

func (user User) validatePhoneNumber() error {
	validationError := customerror.NewValidationError()
	if len(user.PhoneNumber) < phoneNumberMinLen || len(user.PhoneNumber) > phoneNumberMaxLen {
		validationError.AddError("phoneNumber", fmt.Sprintf("must be between %d and %d characters in length", phoneNumberMinLen, phoneNumberMaxLen))
	}

	// check that have prefix +62 and only containt 0-9 after that
	phoneNumberRegex := regexp.MustCompile(`^\+62[0-9]`)
	match := phoneNumberRegex.MatchString(user.PhoneNumber)
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

func (user User) validatePassword() error {
	validationError := customerror.NewValidationError()
	if len(user.Password) < passwordMinLen || len(user.Password) > passwordMaxLen {
		validationError.AddError("password", fmt.Sprintf("must be between %d and %d characters in length", passwordMinLen, passwordMaxLen))
	}

	// Check for at least 1 capital letter, 1 number, and 1 special character
	passwordRegex := regexp.MustCompile(`^(.*[A-Z])(.*\d)(.*[^A-Za-z0-9])`)
	match := passwordRegex.MatchString(user.Password)
	if !match {
		validationError.AddError("password", "containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters")
	}

	if validationError.HasError() {
		return validationError
	}

	return nil
}
