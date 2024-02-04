package driven

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driven"
)

var _ driven.UserWriter = new(UserDB)

type UserDB struct {
	conn *PostgreConnection
}

func NewUserDB(db *PostgreConnection) *UserDB {
	return &UserDB{
		conn: db,
	}
}

// Create implements driven.UserWriter.
func (udb *UserDB) Create(ctx context.Context, user *entity.User) (id string, err error) {
	err = udb.conn.Db.QueryRowContext(ctx, `
		INSERT INTO
			users (full_name, phone_number, password)
		VALUES
			($1, $2, $3)
		RETURNING
			id
	`, user.FullName, user.PhoneNumber, user.Password).Scan(&id)
	return
}
