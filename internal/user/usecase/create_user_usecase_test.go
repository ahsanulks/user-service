package usecase_test

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/param/request"
	"github.com/SawitProRecruitment/UserService/internal/user/param/response"
	"github.com/SawitProRecruitment/UserService/internal/user/usecase"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/test/fake"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	fakeUserDriven := fake.NewFakeUserDriven()
	bcrypt := new(repository.BcyrpEncryption)
	type args struct {
		ctx   context.Context
		param *request.CreateUser
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "when phoneNumber less than 10 character it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+62123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 13 characters in length",
		},
		{
			name: "when phoneNumber more than 13 character it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+62123123123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 13 characters in length",
		},
		{
			name: "when phoneNumber not have prefix +62 it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "0812311231231",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must start with '+62' and only containt number",
		},
		{
			name: "when phoneNumber have prefix +62 but containt other than number, it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+62abc3123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must start with '+62' and only containt number",
		},
		{
			name: "when fullName less than 3 character, it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+628123123123",
				FullName:    generateRandomString(2, ""),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "fullName: must be between 3 and 60 characters in length",
		},
		{
			name: "when fullName more than 60 character, it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+628123123123",
				FullName:    generateRandomString(61, ""),
				Password:    generateRandomPassword(12),
			}},
			wantErr:    true,
			wantErrMsg: "fullName: must be between 3 and 60 characters in length",
		},
		{
			name: "when password less than 6 character, it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "aA2.",
			}},
			wantErr:    true,
			wantErrMsg: "password: must be between 6 and 64 characters in length",
		},
		{
			name: "when password more than 64 character, it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(65),
			}},
			wantErr:    true,
			wantErrMsg: "password: must be between 6 and 64 characters in length",
		},
		{
			name: "when password not containt number, it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "abcAdasd.D",
			}},
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when password not containt capital char, it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "abc123.asd",
			}},
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when password not containt special char, it should return error",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    "Asd123ASd",
			}},
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when all field not valid, it should return all error message",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "",
				FullName:    "",
				Password:    "",
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 13 characters in length,must start with '+62' and only containt number;fullName: must be between 3 and 60 characters in length;password: must be between 6 and 64 characters in length,containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when all field valid, it should return id and saved",
			args: args{context.Background(), &request.CreateUser{
				PhoneNumber: "+628123123123",
				FullName:    faker.Name(),
				Password:    generateRandomPassword(12),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := usecase.NewUserUsecase(fakeUserDriven, bcrypt, nil, nil)
			gotID, err := uu.CreateUser(tt.args.ctx, tt.args.param)
			assert := assert.New(t)
			if tt.wantErr {
				assert.Error(err)
				assert.True(assertMessagesEqual(tt.wantErrMsg, err.Error()))
			} else {
				gotUser, err := fakeUserDriven.GetByID(tt.args.ctx, gotID)
				assert.NoError(err)
				assert.Equal(gotUser.ID, gotID)
			}
		})
	}
}

func splitAndSortErrorMessage(message string) []string {
	slicedMessages := strings.Split(message, ";")
	for i, msg := range slicedMessages {
		slicedMessages[i] = strings.TrimSpace(msg)
	}

	sort.Strings(slicedMessages)
	return slicedMessages
}

func assertMessagesEqual(message1, message2 string) bool {
	sorted1 := splitAndSortErrorMessage(message1)
	sorted2 := splitAndSortErrorMessage(message2)

	return fmt.Sprintf("%v", sorted1) == fmt.Sprintf("%v", sorted2)
}

func TestCreateUser_withPasswordEncrypted(t *testing.T) {
	fakeUserDriven := fake.NewFakeUserDriven()
	bcrypt := new(repository.BcyrpEncryption)
	uu := usecase.NewUserUsecase(fakeUserDriven, bcrypt, nil, nil)
	assert := assert.New(t)

	userParam := &request.CreateUser{
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

func TestUserUsecase_UpdateProfileByID(t *testing.T) {
	fakeUserDriven := fake.NewFakeUserDriven()
	user := &entity.User{
		PhoneNumber: "+628123123123",
		FullName:    faker.Name(),
	}
	id, _ := fakeUserDriven.Create(context.Background(), user)
	user.ID = id

	invalidName := "as"
	invalidPhone := "09123131"

	type args struct {
		ctx    context.Context
		id     string
		params *request.UpdateProfile
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "when all params empy it should return error",
			args: args{
				context.Background(),
				faker.UUIDHyphenated(),
				&request.UpdateProfile{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when fullName not match with criteria it should return error",
			args: args{
				context.Background(),
				faker.UUIDHyphenated(),
				&request.UpdateProfile{
					FullName: &invalidName,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when phoneNumber not match with criteria it should return error",
			args: args{
				context.Background(),
				faker.UUIDHyphenated(),
				&request.UpdateProfile{
					PhoneNumber: &invalidPhone,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when got error from writer it should return error",
			args: args{
				context.Background(),
				"1231313121",
				&request.UpdateProfile{
					FullName:    &user.FullName,
					PhoneNumber: &user.PhoneNumber,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when success it should return user",
			args: args{
				context.Background(),
				user.ID,
				&request.UpdateProfile{
					FullName:    &user.FullName,
					PhoneNumber: &user.PhoneNumber,
				},
			},
			want:    user,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := usecase.NewUserUsecase(fakeUserDriven, nil, nil, nil)
			result, err := uu.UpdateProfileByID(tt.args.ctx, tt.args.id, tt.args.params)
			assert := assert.New(t)
			if tt.wantErr {
				assert.Error(err)
			} else {
				gotUser, err := fakeUserDriven.GetByID(tt.args.ctx, result.ID)
				assert.NoError(err)
				assert.Equal(gotUser.FullName, result.FullName)
				assert.Equal(gotUser.PhoneNumber, result.PhoneNumber)
			}
		})
	}
}

func TestUserUsecase_GenerateUserToken(t *testing.T) {
	fakeUserDriven := fake.NewFakeUserDriven()
	validPassword := faker.Password()
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(validPassword), bcrypt.DefaultCost)
	user := &entity.User{
		PhoneNumber: "+628123123123",
		FullName:    faker.Name(),
		Password:    string(encryptedPassword),
	}
	fakeUserDriven.Create(context.Background(), user)

	invalidUser := &entity.User{
		PhoneNumber: "======111",
		FullName:    faker.Name(),
		Password:    faker.Password(),
	}
	fakeUserDriven.Create(context.Background(), invalidUser)

	type args struct {
		ctx    context.Context
		params *request.GenerateUserTokenRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *response.Token
		wantErr bool
	}{
		{
			name: "when user not found, it should return error",
			args: args{
				context.Background(),
				&request.GenerateUserTokenRequest{
					PhoneNumber: faker.Phonenumber(),
					Password:    faker.Password(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when user password error, it should return error",
			args: args{
				context.Background(),
				&request.GenerateUserTokenRequest{
					PhoneNumber: user.Password,
					Password:    faker.Password(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when generate token error, it should return error",
			args: args{
				context.Background(),
				&request.GenerateUserTokenRequest{
					PhoneNumber: invalidUser.PhoneNumber,
					Password:    invalidUser.Password,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when error record last login count, it should return token",
			args: args{
				context.WithValue(context.Background(), "token_error", true),
				&request.GenerateUserTokenRequest{
					PhoneNumber: user.PhoneNumber,
					Password:    validPassword,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success, it should return token",
			args: args{
				context.Background(),
				&request.GenerateUserTokenRequest{
					PhoneNumber: user.PhoneNumber,
					Password:    validPassword,
				},
			},
			want: &response.Token{
				Token:     "1231313213213131",
				ExpiresIn: 3600,
				Type:      "Bearer",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := usecase.NewUserUsecase(fakeUserDriven, new(repository.BcyrpEncryption), fakeUserDriven, new(fake.FakeTokenProvider))
			result, err := uu.GenerateUserToken(tt.args.ctx, tt.args.params)
			assert := assert.New(t)
			assert.Equal(tt.wantErr, err != nil)
			assert.Equal(tt.want, result)
		})
	}
}
