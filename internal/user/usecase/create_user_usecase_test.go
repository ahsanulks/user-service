package usecase_test

import (
	"context"
	"crypto/rand"
	"math/big"
	"regexp"
	"testing"

	"github.com/SawitProRecruitment/UserService/adapter/driven"
	"github.com/SawitProRecruitment/UserService/internal/user/usecase"
	"github.com/SawitProRecruitment/UserService/test/fake"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	fakeUserDriven := fake.NewFakeUserDriven()
	bcrypt := new(driven.BcyrpEncryption)
	type args struct {
		ctx   context.Context
		param *usecase.CreateUserParam
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "when phoneNumber less than 10 character it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+62123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 13 characters in length",
		},
		{
			name: "when phoneNumber more than 13 character it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+62123123123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 13 characters in length",
		},
		{
			name: "when phoneNumber not have prefix +62 it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "0812311231231",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must start with '+62' and only containt number",
		},
		{
			name: "when phoneNumber have prefix +62 but containt other than number, it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+62abc3123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must start with '+62' and only containt number",
		},
		{
			name: "when fullName less than 3 character, it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    generateRandomString(2, ""),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "fullName: must be between 3 and 60 characters in length",
		},
		{
			name: "when fullName more than 60 character, it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    generateRandomString(61, ""),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "fullName: must be between 3 and 60 characters in length",
		},
		{
			name: "when password less than 6 character, it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "aA2.",
			}},
			wantErr:    true,
			wantErrMsg: "password: must be between 6 and 64 characters in length",
		},
		{
			name: "when password more than 64 character, it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(65),
			}},
			wantErr:    true,
			wantErrMsg: "password: must be between 6 and 64 characters in length",
		},
		{
			name: "when password not containt number, it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "abcAdasd.D",
			}},
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when password not containt capital char, it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "abc123.asd",
			}},
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when password not containt special char, it should return error",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "Asd123ASd",
			}},
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when all field not valid, it should return all error message",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "",
				FullName:    "",
				Password:    "",
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 13 characters in length,must start with '+62' and only containt number;fullName: must be between 3 and 60 characters in length;password: must be between 6 and 64 characters in length,containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when all field valid, it should return id and saved",
			args: args{context.Background(), &usecase.CreateUserParam{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := usecase.NewUserUsecase(fakeUserDriven, bcrypt)
			gotID, err := uu.CreateUser(tt.args.ctx, tt.args.param)
			assert := assert.New(t)
			if tt.wantErr {
				assert.Error(err)
				assert.EqualError(err, tt.wantErrMsg)
			} else {
				gotUser, err := fakeUserDriven.GetByID(tt.args.ctx, gotID)
				assert.NoError(err)
				assert.Equal(gotUser.ID, gotID)
			}
		})
	}
}

func TestCreateUser_withPasswordEncrypted(t *testing.T) {
	fakeUserDriven := fake.NewFakeUserDriven()
	bcrypt := new(driven.BcyrpEncryption)
	uu := usecase.NewUserUsecase(fakeUserDriven, bcrypt)
	assert := assert.New(t)

	userParam := &usecase.CreateUserParam{
		PhoneNumber: "+628123123123",
		FullName:    faker.Name(),
		Password:    generateRandomPassword(12),
	}
	gotID, err := uu.CreateUser(context.Background(), userParam)
	assert.NoError(err)

	user, err := fakeUserDriven.GetByID(context.Background(), gotID)
	assert.NoError(err)
	assert.Equal(user.ID, gotID)
	assert.NotEqual(userParam.Password, user.Password)

	err = bcrypt.CompareEncryptedAndData([]byte(user.Password), []byte(userParam.Password))
	assert.NoError(err)
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
