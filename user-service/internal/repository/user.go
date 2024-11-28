package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type User struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUser(db *gorm.DB, redis *redis.Client) *User {
	return &User{db: db, redis: redis}
}

func (u *User) GetByID(ctx context.Context, id uint) (*model.User, error) {
	user := &model.User{}
	userKey := fmt.Sprintf("user:%d", id)
	userData, err := u.redis.Get(ctx, userKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(userData), user)
		if err != nil {
			return nil, utils.NewError(utils.ErrInternalServer, err.Error())
		}
		return user, nil
	}
	err = u.db.WithContext(ctx).First(user, id).Error
	if err != nil {
		return nil, utils.NewError(utils.ErrNotFound, err.Error())
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
		return utils.NewError(utils.ErrInternalServer, err.Error())
	}
	return nil
}

func (u *User) Update(ctx context.Context, id uint, user *model.User) error {
	err := u.db.WithContext(ctx).Model(user).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return utils.NewError(utils.ErrInternalServer, err.Error())
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
		return utils.NewError(utils.ErrInternalServer, err.Error())
	}
	userKey := fmt.Sprintf("user:%d", id)
	go func() {
		u.redis.Del(ctx, userKey)
	}()
	return nil
}
