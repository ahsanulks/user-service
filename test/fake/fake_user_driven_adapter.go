package fake

import (
	"context"
	"errors"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driven"
	"github.com/google/uuid"
)

var _ driven.UserWriter = new(FakeUserDriven)

type FakeUserDriven struct {
	data map[string]*entity.User
}

func NewFakeUserDriven() *FakeUserDriven {
	return &FakeUserDriven{
		data: make(map[string]*entity.User),
	}
}

// Create implements driven.UserWriter.
func (fud *FakeUserDriven) Create(ctx context.Context, user *entity.User) (id string, err error) {
	user.ID = uuid.New().String()
	fud.data[user.ID] = user
	return user.ID, nil
}

func (fud FakeUserDriven) GetByID(ctx context.Context, id string) (*entity.User, error) {
	if user, ok := fud.data[id]; ok {
		return user, nil
	}
	return nil, errors.New("resource not found")
}