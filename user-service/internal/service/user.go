package service

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/repository"
)

type User struct {
	repoUser *repository.User
}

func NewUser(repoUser *repository.User) *User {
	return &User{repoUser: repoUser}
}

func (u *User) GetByID(ctx context.Context, id uint) (*model.User, error) {
	return u.repoUser.GetByID(ctx, id)
}

func (u *User) Create(ctx context.Context, user *model.User) error {
	return u.repoUser.Create(ctx, user)
}

func (u *User) Update(ctx context.Context, id uint, user *model.User) error {
	return u.repoUser.Update(ctx, id, user)
}

func (u *User) Delete(ctx context.Context, id uint) error {
	return u.repoUser.Delete(ctx, id)
}
