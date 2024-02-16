package usecase_test

import (
	"context"
	"testing"

	"userservice/internal/user/entity"
	"userservice/internal/user/usecase"
	"userservice/test/fake"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_GetUserByID(t *testing.T) {
	fakeUserDriven := fake.NewFakeUserDriven()
	user := &entity.User{
		FullName:    faker.Name(),
		PhoneNumber: faker.Phonenumber(),
		Password:    faker.Password(),
	}
	id, _ := fakeUserDriven.Create(context.Background(), user)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "when user found it should return entity",
			args: args{
				context.Background(),
				id,
			},
			want: &entity.User{
				ID:          id,
				FullName:    user.FullName,
				PhoneNumber: user.PhoneNumber,
			},
			wantErr: false,
		},
		{
			name: "when user not found it should return error",
			args: args{
				context.Background(),
				"1231331",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := usecase.NewUserGetterUsecase(fakeUserDriven)
			got, err := uu.GetUserByID(tt.args.ctx, tt.args.id)
			assert := assert.New(t)

			assert.Equal(tt.wantErr, err != nil)
			if tt.wantErr {
				assert.Nil(got)
			} else {
				assert.Equal(tt.want.ID, got.ID)
				assert.Equal(tt.want.FullName, got.FullName)
				assert.Equal(tt.want.PhoneNumber, got.PhoneNumber)
			}
		})
	}
}
