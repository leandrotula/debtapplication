package storage

import (
	"context"
	model2 "debtsapp/internal/service/model"
	"debtsapp/internal/storage/model"
	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (u *UserMock) CreateAndInvite(ctx context.Context, user *model2.UserRequest, token string) error {
	args := u.Called(ctx, user, token)
	return args.Error(0)
}

func (u *UserMock) Activate(ctx context.Context, token string) error {
	args := u.Called(ctx, token)
	return args.Error(0)
}

func (u *UserMock) FindUserByEmail(ctx context.Context, email string) (*model.UserEntity, error) {
	args := u.Called(ctx, email)
	user := args.Get(0)
	err := args.Error(1)
	if user == nil {
		return nil, err
	}
	return user.(*model.UserEntity), args.Error(1)
}
