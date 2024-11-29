package repository

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/author-service/internal/model"
	"gorm.io/gorm"
)

type Author struct {
	db *gorm.DB
}

type AuthorRepository interface {
	Create(ctx context.Context, author *model.Author) error
	ReadAll(ctx context.Context) ([]model.Author, error)
	ReadByID(ctx context.Context, id int) (model.Author, error)
	Update(ctx context.Context, author *model.Author) error
	Delete(ctx context.Context, id int) error
}

func NewAuthor(db *gorm.DB) AuthorRepository {
	return &Author{db: db}
}

func (a *Author) Create(ctx context.Context, author *model.Author) error {
	return a.db.WithContext(ctx).Create(author).Error
}

func (a *Author) ReadAll(ctx context.Context) ([]model.Author, error) {
	var authors []model.Author
	err := a.db.WithContext(ctx).Find(&authors).Error
	return authors, err
}

func (a *Author) ReadByID(ctx context.Context, id int) (model.Author, error) {
	var author model.Author
	err := a.db.WithContext(ctx).First(&author, id).Error
	return author, err
}

func (a *Author) Update(ctx context.Context, author *model.Author) error {
	return a.db.WithContext(ctx).Save(author).Error
}

func (a *Author) Delete(ctx context.Context, id int) error {
	return a.db.WithContext(ctx).Delete(&model.Author{}, id).Error
}
