package service

import (
	"context"
	"strconv"

	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/repository"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/utils"
)

type User struct {
	repoUser *repository.User
}

type UserService interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, id string, user *model.User) error
	Delete(ctx context.Context, id string) error
}

func NewUser(repoUser *repository.User) UserService {
	return &User{repoUser: repoUser}
}

func (u *User) GetByID(ctx context.Context, id string) (*model.User, error) {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, err.Error())
	}
	return u.repoUser.GetByID(ctx, uint(idUint))
}

func (u *User) Create(ctx context.Context, user *model.User) error {
	return u.repoUser.Create(ctx, user)
}

func (u *User) Update(ctx context.Context, id string, user *model.User) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return utils.NewError(utils.ErrBadRequest, err.Error())
	}
	return u.repoUser.Update(ctx, uint(idUint), user)
}

func (u *User) Delete(ctx context.Context, id string) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return utils.NewError(utils.ErrBadRequest, err.Error())
	}
	return u.repoUser.Delete(ctx, uint(idUint))
}
