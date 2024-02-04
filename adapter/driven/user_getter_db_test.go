package driven

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
