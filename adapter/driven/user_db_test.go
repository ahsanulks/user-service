package driven

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/param/request"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func newMockConn() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestUserDB_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *entity.User
	}
	tests := []struct {
		name       string
		args       args
		wantId     string
		wantErr    bool
		expectFunc func(sqlmock.Sqlmock, *entity.User)
	}{
		{
			name: "when given user entity, it should insert fullName, phoneNumber, password",
			args: args{
				ctx: context.Background(),
				user: &entity.User{
					FullName:    faker.Name(),
					PhoneNumber: faker.Phonenumber(),
					Password:    faker.Password(),
				},
			},
			wantId:  "aaaa-bbb-ccc-123",
			wantErr: false,
			expectFunc: func(mock sqlmock.Sqlmock, user *entity.User) {
				mock.ExpectQuery("^INSERT INTO users").
					WithArgs(user.FullName, user.PhoneNumber, user.Password).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("aaaa-bbb-ccc-123"))
			},
		},
		{
			name: "when something went wrong in db, it should return error",
			args: args{
				ctx: context.Background(),
				user: &entity.User{
					FullName:    faker.Name(),
					PhoneNumber: faker.Phonenumber(),
					Password:    faker.Password(),
				},
			},
			wantId:  "",
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, user *entity.User) {
				mock.ExpectQuery("^INSERT INTO users").
					WithArgs(user.FullName, user.PhoneNumber, user.Password).
					WillReturnError(errors.New("some database error"))
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

			tt.expectFunc(dbMock, tt.args.user)

			gotId, err := udb.Create(tt.args.ctx, tt.args.user)

			assert := assert.New(t)
			assert.Equal(tt.wantErr, err != nil)
			assert.Equal(tt.wantId, gotId)
			assert.NoError(dbMock.ExpectationsWereMet())
		})
	}
}

func TestUserDB_UpdateProfileByID(t *testing.T) {
	validUser := &entity.User{
		ID:          faker.UUIDHyphenated(),
		FullName:    faker.Name(),
		PhoneNumber: faker.Phonenumber(),
	}
	type args struct {
		ctx    context.Context
		id     string
		params *request.UpdateProfile
	}
	tests := []struct {
		name       string
		args       args
		want       *entity.User
		wantErr    bool
		expectFunc func(sqlmock.Sqlmock, *entity.User)
	}{
		{
			name: "when user not found, should return error",
			args: args{
				context.Background(),
				validUser.ID,
				&request.UpdateProfile{
					FullName: ptrString(validUser.FullName),
				},
			},
			want:    nil,
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, _ *entity.User) {
				mock.ExpectQuery("UPDATE users SET").WillReturnError(sql.ErrNoRows)

			},
		},
		{
			name: "when containt duplicate, should return error",
			args: args{
				context.Background(),
				validUser.ID,
				&request.UpdateProfile{
					FullName: ptrString(validUser.FullName),
				},
			},
			want:    nil,
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, _ *entity.User) {
				mock.ExpectQuery("UPDATE users SET").
					WillReturnError(errors.New("pq: duplicate key value violates unique constraint"))
			},
		},
		{
			name: "when only update full_name, should return user",
			args: args{
				context.Background(),
				validUser.ID,
				&request.UpdateProfile{
					FullName: ptrString(validUser.FullName),
				},
			},
			want:    validUser,
			wantErr: false,
			expectFunc: func(mock sqlmock.Sqlmock, expectedUser *entity.User) {
				mock.ExpectQuery("UPDATE users SET").
					WithArgs(validUser.FullName, validUser.ID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "full_name", "phone_number"}).
						AddRow(expectedUser.ID, expectedUser.FullName, expectedUser.PhoneNumber))
			},
		},
		{
			name: "when only update phone_number, should return user",
			args: args{
				context.Background(),
				validUser.ID,
				&request.UpdateProfile{
					PhoneNumber: ptrString(validUser.PhoneNumber),
				},
			},
			want:    validUser,
			wantErr: false,
			expectFunc: func(mock sqlmock.Sqlmock, expectedUser *entity.User) {
				mock.ExpectQuery("UPDATE users SET").
					WithArgs(validUser.PhoneNumber, validUser.ID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "full_name", "phone_number"}).
						AddRow(expectedUser.ID, expectedUser.FullName, expectedUser.PhoneNumber))
			},
		},
		{
			name: "when update phone_number and full_name, should return user",
			args: args{
				context.Background(),
				validUser.ID,
				&request.UpdateProfile{
					PhoneNumber: ptrString(validUser.PhoneNumber),
					FullName:    ptrString(validUser.FullName),
				},
			},
			want:    validUser,
			wantErr: false,
			expectFunc: func(mock sqlmock.Sqlmock, expectedUser *entity.User) {
				mock.ExpectQuery("UPDATE users SET").
					WithArgs(validUser.FullName, validUser.PhoneNumber, validUser.ID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "full_name", "phone_number"}).
						AddRow(expectedUser.ID, expectedUser.FullName, expectedUser.PhoneNumber))
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

			result, err := udb.UpdateProfileByID(tt.args.ctx, tt.args.id, tt.args.params)

			assert := assert.New(t)
			assert.Equal(tt.wantErr, err != nil)
			assert.Equal(tt.want, result)
			assert.NoError(dbMock.ExpectationsWereMet())
		})
	}
}

func ptrString(s string) *string {
	return &s
}

func TestUserDB_UpdateUserToken(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		expectFunc func(sqlmock.Sqlmock, string)
	}{
		{
			name: "when error on db, it should return error",
			args: args{
				context.Background(),
				"12313",
			},
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, userId string) {
				mock.ExpectExec("INSERT INTO user_tokens").
					WithArgs(userId).
					WillReturnError(errors.New("some database error"))
			},
		},
		{
			name: "when error on db, it should return error",
			args: args{
				context.Background(),
				"12313",
			},
			wantErr: false,
			expectFunc: func(mock sqlmock.Sqlmock, userId string) {
				mock.ExpectExec("INSERT INTO user_tokens").
					WithArgs(userId).
					WillReturnResult(sqlmock.NewResult(0, 1))
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

			tt.expectFunc(dbMock, tt.args.userId)

			err := udb.UpdateUserToken(tt.args.ctx, tt.args.userId)

			assert := assert.New(t)
			assert.Equal(tt.wantErr, err != nil)
			assert.NoError(dbMock.ExpectationsWereMet())
		})
	}
}
