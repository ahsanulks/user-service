package driven

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/internal/user/entity"
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
