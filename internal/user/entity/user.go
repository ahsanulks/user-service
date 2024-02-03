package entity

import (
	"fmt"
	"regexp"
	"time"
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
		ID:          "",
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
	if len(user.PhoneNumber) < phoneNumberMinLen || len(user.PhoneNumber) > phoneNumberMaxLen {
		return fmt.Errorf("phone number must be between %d and %d characters in length", phoneNumberMinLen, phoneNumberMaxLen)
	}

	phoneNumberRegex := `^\+62[0-9]$`
	match, _ := regexp.MatchString(phoneNumberRegex, user.PhoneNumber)
	if !match {
		return fmt.Errorf("phone number must start with '+62' and only containt number")
	}
	return nil
}
