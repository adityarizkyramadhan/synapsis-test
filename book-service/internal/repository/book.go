package repository

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"gorm.io/gorm"
)

type BookService interface {
	Create(ctx context.Context, book *model.Book) (*model.Book, error)
	GetAll(ctx context.Context) ([]model.Book, error)
	GetByID(ctx context.Context, id uint32) (*model.Book, error)
	Update(ctx context.Context, book *model.Book) (*model.Book, error)
	Delete(ctx context.Context, id uint32) error
}

type book struct {
	db *gorm.DB
}

func Newbook(db *gorm.DB) BookService {
	return &book{db}
}

func (c *book) Create(ctx context.Context, book *model.Book) (*model.Book, error) {
	if err := c.db.WithContext(ctx).Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (c *book) GetAll(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	if err := c.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (c *book) GetByID(ctx context.Context, id uint32) (*model.Book, error) {
	var book model.Book
	if err := c.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (c *book) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	if err := c.db.WithContext(ctx).Save(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (c *book) Delete(ctx context.Context, id uint32) error {
	if err := c.db.WithContext(ctx).Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
