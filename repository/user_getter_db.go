package repository

import (
	"context"
	"database/sql"

	"userservice/internal/user/entity"

	"github.com/google/uuid"
)

// GetByID implements driven.UserGetter.
func (udb *UserDB) GetByID(ctx context.Context, id string) (*entity.User, error) {
	uuidUser, _ := uuid.Parse(id)
	rows, err := udb.conn.Db.QueryContext(ctx, `
		SELECT
			id,
			full_name,
			phone_number,
			password,
			created_at,
			updated_at
		FROM
			users
		WHERE
			id = $1
		LIMIT
			1
	`, uuidUser)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var user entity.User
	if rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.FullName,
			&user.PhoneNumber,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	} else {
		return nil, sql.ErrNoRows
	}

	return &user, err
}

// GetByPhoneNumber implements driven.UserGetter.
func (udb *UserDB) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	rows, err := udb.conn.Db.QueryContext(ctx, `
		SELECT
			id,
			full_name,
			phone_number,
			password,
			created_at,
			updated_at
		FROM
			users
		WHERE
			phone_number = $1
		LIMIT
			1
	`, phoneNumber)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var user entity.User
	if rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.FullName,
			&user.PhoneNumber,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	} else {
		return nil, sql.ErrNoRows
	}

	return &user, err
}
