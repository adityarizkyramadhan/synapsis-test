package repository

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"gorm.io/gorm"
)

type CategoryBookRepository interface {
	Add(ctx context.Context, category *model.Book) (*model.Book, error)
	Delete(ctx context.Context, id uint32) error
	GetAll(ctx context.Context) ([]model.Book, error)
}

type categoryBook struct {
	db *gorm.DB
}

func NewCategoryBook(db *gorm.DB) CategoryBookRepository {
	return &categoryBook{db}
}

func (c *categoryBook) Add(ctx context.Context, category *model.Book) (*model.Book, error) {
	if err := c.db.WithContext(ctx).Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryBook) GetAll(ctx context.Context) ([]model.Book, error) {
	var categories []model.Book
	if err := c.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *categoryBook) Delete(ctx context.Context, id uint32) error {
	if err := c.db.WithContext(ctx).Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
