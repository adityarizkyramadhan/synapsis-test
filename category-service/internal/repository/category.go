package repository

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/category-service/internal/model"
	"gorm.io/gorm"
)

type Category struct {
	db *gorm.DB
}

type CategoryRepository interface {
	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id uint32) error
	GetByID(ctx context.Context, id uint32) (*model.Category, error)
	GetAll(ctx context.Context) ([]model.Category, error)
}

func NewCategory(db *gorm.DB) CategoryRepository {
	return &Category{db}
}

func (c *Category) Create(ctx context.Context, category *model.Category) error {
	if err := c.db.WithContext(ctx).Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) Update(ctx context.Context, category *model.Category) error {
	if err := c.db.WithContext(ctx).Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) Delete(ctx context.Context, id uint32) error {
	if err := c.db.WithContext(ctx).Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) GetByID(ctx context.Context, id uint32) (*model.Category, error) {
	var category model.Category
	if err := c.db.WithContext(ctx).First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *Category) GetAll(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	if err := c.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
