package usecase

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	faker.Username()
	type args struct {
		ctx   context.Context
		param *CreateUserParam
	}
	tests := []struct {
		name    string
		uu      *UserUsecase
		args    args
		wantId  string
		wantErr bool
	}{
		{
			name: "when phoneNumber less than 10 character it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+62123",
				FullName:    faker.Name(options.WithRandomStringLength(10)),
				Password:    faker.Password(options.WithRandomStringLength(10)),
			}},
			wantId:  "",
			wantErr: true,
		},
		{
			name: "when phoneNumber more than 13 character it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+62123123123123",
				FullName:    faker.Name(options.WithRandomStringLength(10)),
				Password:    faker.Password(options.WithRandomStringLength(10)),
			}},
			wantId:  "",
			wantErr: true,
		},
		{
			name: "when phoneNumber not have prefix +62 it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "0812311231231",
				FullName:    faker.Name(options.WithRandomStringLength(10)),
				Password:    faker.Password(options.WithRandomStringLength(10)),
			}},
			wantId:  "",
			wantErr: true,
		},
		{
			name: "when phoneNumber have prefix +62 but containt other than number, it should return error",
			uu:   NewUserUsecase(),
			args: args{context.Background(), &CreateUserParam{
				PhoneNumber: "+62abc3123123",
				FullName:    faker.Name(options.WithRandomStringLength(10)),
				Password:    faker.Password(options.WithRandomStringLength(10)),
			}},
			wantId:  "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := tt.uu.CreateUser(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("UserUsecase.CreateUser() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestCreateUser_shouldReturnErrorAllConstraintValidation(t *testing.T) {
	userUsecase := NewUserUsecase()
	_, err := userUsecase.CreateUser(context.Background(), &CreateUserParam{
		PhoneNumber: "123131",
		FullName:    faker.Name(options.WithRandomStringLength(10)),
		Password:    faker.Password(options.WithRandomStringLength(10)),
	})
	assert := assert.New(t)
	assert.Error(err)
	assert.EqualError(err, "phoneNumber: must be between 10 and 13 characters in length;must start with '+62' and only containt number")
}
