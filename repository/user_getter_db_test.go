package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserDB_GetByID(t *testing.T) {
	validUserID := faker.UUIDHyphenated()
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name       string
		args       args
		want       *entity.User
		wantErr    bool
		expectFunc func(sqlmock.Sqlmock, *entity.User)
	}{
		{
			name: "when record not found, it should return error",
			args: args{
				context.Background(),
				faker.UUIDHyphenated(),
			},
			want:    nil,
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, _ *entity.User) {
				mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{}))
			},
		},
		{
			name: "when error on database, it should return error",
			args: args{
				context.Background(),
				faker.UUIDHyphenated(),
			},
			want:    nil,
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, _ *entity.User) {
				mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("database error"))
			},
		},
		{
			name: "when user found, it should return user",
			args: args{
				context.Background(),
				validUserID,
			},
			want: &entity.User{
				ID:          validUserID,
				FullName:    faker.Name(),
				PhoneNumber: faker.Phonenumber(),
				Password:    faker.Password(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantErr: false,
			expectFunc: func(mock sqlmock.Sqlmock, expectedUser *entity.User) {
				rows := sqlmock.NewRows([]string{"id", "full_name", "phone_number", "password", "created_at", "updated_at"}).
					AddRow(expectedUser.ID, expectedUser.FullName, expectedUser.PhoneNumber, expectedUser.Password, expectedUser.CreatedAt, expectedUser.UpdatedAt)

				mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		conn, dbMock := newMockConn()
		defer conn.Close()
		pgConn := PostgreConnection{
			Db: conn,
		}
		udb := &UserDB{
			conn: &pgConn,
		}

		tt.expectFunc(dbMock, tt.want)

		got, err := udb.GetByID(tt.args.ctx, tt.args.id)

		assert := assert.New(t)
		assert.Equal(tt.wantErr, err != nil)
		assert.Equal(tt.want, got)
		assert.NoError(dbMock.ExpectationsWereMet())
	}
}

func TestUserDB_GetByPhoneNumber(t *testing.T) {
	validPhoneNumber := faker.Phonenumber()
	type args struct {
		ctx         context.Context
		phoneNumber string
	}
	tests := []struct {
		name       string
		args       args
		want       *entity.User
		wantErr    bool
		expectFunc func(sqlmock.Sqlmock, *entity.User)
	}{
		{
			name: "when record not found, it should return error",
			args: args{
				context.Background(),
				faker.Phonenumber(),
			},
			want:    nil,
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, _ *entity.User) {
				mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{}))
			},
		},
		{
			name: "when error on database, it should return error",
			args: args{
				context.Background(),
				faker.Phonenumber(),
			},
			want:    nil,
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, _ *entity.User) {
				mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("database error"))
			},
		},
		{
			name: "when user found, it should return user",
			args: args{
				context.Background(),
				validPhoneNumber,
			},
			want: &entity.User{
				ID:          faker.UUIDHyphenated(),
				FullName:    faker.Name(),
				PhoneNumber: validPhoneNumber,
				Password:    faker.Password(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantErr: false,
			expectFunc: func(mock sqlmock.Sqlmock, expectedUser *entity.User) {
				rows := sqlmock.NewRows([]string{"id", "full_name", "phone_number", "password", "created_at", "updated_at"}).
					AddRow(expectedUser.ID, expectedUser.FullName, expectedUser.PhoneNumber, expectedUser.Password, expectedUser.CreatedAt, expectedUser.UpdatedAt)

				mock.ExpectQuery("SELECT").WithArgs(validPhoneNumber).WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, dbMock := newMockConn()
			defer conn.Close()
			pgConn := PostgreConnection{
				Db: conn,
			}
			udb := &UserDB{
				conn: &pgConn,
			}

			tt.expectFunc(dbMock, tt.want)

			got, err := udb.GetByPhoneNumber(tt.args.ctx, tt.args.phoneNumber)

			assert := assert.New(t)
			assert.Equal(tt.wantErr, err != nil)
			assert.Equal(tt.want, got)
			assert.NoError(dbMock.ExpectationsWereMet())
		})
	}
}
