package usecase

import (
	"context"
	"crypto/rand"
	"math/big"
	"regexp"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	faker.Username()
	type args struct {
		ctx   context.Context
		param *CreateUserParam
	}
	tests := []struct {
		name       string
		uu         *UserUsecase
		args       args
		wantId     string
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "when phoneNumber less than 10 character it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+62123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 13 characters in length",
		},
		{
			name: "when phoneNumber more than 13 character it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+62123123123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 13 characters in length",
		},
		{
			name: "when phoneNumber not have prefix +62 it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "0812311231231",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "phoneNumber: must start with '+62' and only containt number",
		},
		{
			name: "when phoneNumber have prefix +62 but containt other than number, it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+62abc3123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "phoneNumber: must start with '+62' and only containt number",
		},
		{
			name: "when fullName less than 3 character, it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    generateRandomString(2, ""),
				Password:    generateRandomPassword(12),
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "fullName: must be between 3 and 60 characters in length",
		},
		{
			name: "when fullName more than 60 character, it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    generateRandomString(61, ""),
				Password:    generateRandomPassword(12),
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "fullName: must be between 3 and 60 characters in length",
		},
		{
			name: "when password less than 6 character, it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "aA2.",
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "password: must be between 6 and 64 characters in length",
		},
		{
			name: "when password more than 64 character, it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(65),
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "password: must be between 6 and 64 characters in length",
		},
		{
			name: "when password not containt number, it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "abcAdasd.D",
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when password not containt capital char, it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "abc123.asd",
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when password not containt special char, it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "Asd123ASd",
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when all field not valid, it should return all error message",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "",
				FullName:    "",
				Password:    "",
			}},
			wantId:     "",
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 13 characters in length,must start with '+62' and only containt number;fullName: must be between 3 and 60 characters in length;password: must be between 6 and 64 characters in length,containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when all field valid, it should return id and saved",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantId:     "asdad1-1231-asdsa-123123",
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := tt.uu.CreateUser(tt.args.ctx, tt.args.param)
			assert := assert.New(t)
			assert.Equal(tt.wantId, gotId)
			if tt.wantErr {
				assert.Error(err)
				assert.EqualError(err, tt.wantErrMsg)
			}
		})
	}
}

func generateRandomString(length int, charset string) string {
	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	charsetLength := big.NewInt(int64(len(charset)))

	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, charsetLength)
		randomString[i] = charset[randomIndex.Int64()]
	}

	return string(randomString)
}

func generateRandomPassword(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>/?~"
	passwordRegex := regexp.MustCompile(`^(.*[A-Z])(.*\d)(.*[^A-Za-z0-9])`)
	var password string
	match := false
	for !match {
		password = generateRandomString(length, charset)
		match = passwordRegex.MatchString(password)
	}

	return password
}
