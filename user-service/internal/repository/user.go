package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type User struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUser(db *gorm.DB, redis *redis.Client) UserRepository {
	return &User{db: db, redis: redis}
}

type UserRepository interface {
	GetByID(ctx context.Context, id uint) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, id uint, user *model.User) error
	Delete(ctx context.Context, id uint) error
	Login(ctx context.Context, user *model.User) (*model.User, error)
}

func (u *User) Login(ctx context.Context, user *model.User) (*model.User, error) {
	var dataUser *model.User
	err := u.db.WithContext(ctx).Where("email = ?", user.Email).First(dataUser).Error

	if err != nil {
		return nil, err
	}

	return dataUser, nil
}

func (u *User) GetByID(ctx context.Context, id uint) (*model.User, error) {
	user := &model.User{}
	userKey := fmt.Sprintf("user:%d", id)
	userData, err := u.redis.Get(ctx, userKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(userData), user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	err = u.db.WithContext(ctx).First(user, id).Error
	if err != nil {
		return nil, err
	}
	go func() {
		userJSON, _ := json.Marshal(user)
		expired := 3600
		u.redis.Set(ctx, userKey, userJSON, time.Duration(expired)*time.Second)
	}()
	return user, nil
}

func (u *User) Create(ctx context.Context, user *model.User) error {
	err := u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(ctx context.Context, id uint, user *model.User) error {
	err := u.db.WithContext(ctx).Model(user).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return err
	}
	userKey := fmt.Sprintf("user:%d", id)
	go func() {
		u.redis.Del(ctx, userKey)
	}()
	return nil
}

func (u *User) Delete(ctx context.Context, id uint) error {
	err := u.db.WithContext(ctx).Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	userKey := fmt.Sprintf("user:%d", id)
	go func() {
		u.redis.Del(ctx, userKey)
	}()
	return nil
}
