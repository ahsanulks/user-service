package repository

import (
	"context"
	"strconv"
	"strings"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/param/request"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driven"
	"github.com/google/uuid"
)

var (
	_ driven.UserWriter = new(UserDB)
	_ driven.UserGetter = new(UserDB)
)

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

func (udb *UserDB) UpdateProfileByID(ctx context.Context, id string, params *request.UpdateProfile) (*entity.User, error) {
	query := "UPDATE users SET updated_at = now(), "
	var values []any
	var index int

	if params.FullName != nil {
		query += "full_name = $" + strconv.Itoa(index+1) + ", "
		values = append(values, *params.FullName)
		index++
	}

	if params.PhoneNumber != nil {
		query += "phone_number = $" + strconv.Itoa(index+1) + ", "
		values = append(values, *params.PhoneNumber)
		index++
	}

	query = strings.TrimSuffix(query, ", ")

	query += " WHERE id = $" + strconv.Itoa(index+1) + "  RETURNING id, full_name, phone_number"
	uuidUser, _ := uuid.Parse(id)
	values = append(values, uuidUser)

	var user entity.User
	err := udb.conn.Db.QueryRowContext(ctx, query, values...).Scan(
		&user.ID,
		&user.FullName,
		&user.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (udb *UserDB) UpdateUserToken(ctx context.Context, userId string) error {
	_, err := udb.conn.Db.ExecContext(ctx, `
		INSERT INTO
			user_tokens (user_id, success_login_count, last_login_at)
		VALUES
    		($1, 1, NOW())
		ON CONFLICT (user_id)
		DO UPDATE SET
    		success_login_count = user_tokens.success_login_count + 1,
    		last_login_at = NOW()`, userId)
	return err
}
