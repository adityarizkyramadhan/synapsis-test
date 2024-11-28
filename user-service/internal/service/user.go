package service

import (
	"context"
	"strconv"

	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/repository"
)

type User struct {
	repoUser repository.UserRepository
}

type UserService interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, id string, user *model.User) error
	Delete(ctx context.Context, id string) error
}

func NewUser(repoUser repository.UserRepository) UserService {
	return &User{repoUser: repoUser}
}

func (u *User) GetByID(ctx context.Context, id string) (*model.User, error) {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return u.repoUser.GetByID(ctx, uint(idUint))
}

func (u *User) Create(ctx context.Context, user *model.User) error {
	return u.repoUser.Create(ctx, user)
}

func (u *User) Update(ctx context.Context, id string, user *model.User) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return u.repoUser.Update(ctx, uint(idUint), user)
}

func (u *User) Delete(ctx context.Context, id string) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return u.repoUser.Delete(ctx, uint(idUint))
}
