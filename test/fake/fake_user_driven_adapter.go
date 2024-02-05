package fake

import (
	"context"
	"errors"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/param/request"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driven"
	"github.com/google/uuid"
)

var (
	_ driven.UserWriter = new(FakeUserDriven)
	_ driven.UserGetter = new(FakeUserDriven)
)

type FakeUserDriven struct {
	data        map[string]*entity.User
	dataByPhone map[string]*entity.User
}

func NewFakeUserDriven() *FakeUserDriven {
	return &FakeUserDriven{
		data:        make(map[string]*entity.User),
		dataByPhone: make(map[string]*entity.User),
	}
}

// Create implements driven.UserWriter.
func (fud *FakeUserDriven) Create(ctx context.Context, user *entity.User) (id string, err error) {
	user.ID = uuid.New().String()
	// encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// user.Password = string(encryptedPassword)
	fud.data[user.ID] = user
	fud.dataByPhone[user.PhoneNumber] = user
	return user.ID, nil
}

// Create implements driven.UserGetter.
func (fud FakeUserDriven) GetByID(ctx context.Context, id string) (*entity.User, error) {
	if user, ok := fud.data[id]; ok {
		return user, nil
	}
	return nil, errors.New("resource not found")
}

// UpdateProfileByID implements driven.UserWriter.
func (fud *FakeUserDriven) UpdateProfileByID(ctx context.Context, id string, params *request.UpdateProfile) (*entity.User, error) {
	if user, ok := fud.data[id]; ok {
		user.FullName = *params.FullName
		user.PhoneNumber = *params.PhoneNumber
		return user, nil
	}
	return nil, errors.New("resource not found")
}

// GetByPhoneNumber implements driven.UserGetter.
func (fud *FakeUserDriven) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	if user, ok := fud.dataByPhone[phoneNumber]; ok {
		return user, nil
	}
	return nil, errors.New("resource not found")
}

// UpdateUserToken implements driven.UserWriter.
func (*FakeUserDriven) UpdateUserToken(ctx context.Context, userId string) error {
	if val := ctx.Value("token_error"); val != nil {
		return errors.New("error")
	}
	return nil
}
