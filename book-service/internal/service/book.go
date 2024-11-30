package service

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/repository"
)

type BookService interface {
	Create(ctx context.Context, book *model.Book) (*model.Book, error)
	GetAll(ctx context.Context) ([]model.Book, error)
	GetByID(ctx context.Context, id uint32) (*model.Book, error)
	Update(ctx context.Context, book *model.Book) (*model.Book, error)
	Delete(ctx context.Context, id uint32) error
}

type book struct {
	repo repository.BookService
}

func NewBook(repo repository.BookService) BookService {
	return &book{repo}
}

func (c *book) Create(ctx context.Context, book *model.Book) (*model.Book, error) {
	return c.repo.Create(ctx, book)
}

func (c *book) GetAll(ctx context.Context) ([]model.Book, error) {
	return c.repo.GetAll(ctx)
}

func (c *book) GetByID(ctx context.Context, id uint32) (*model.Book, error) {
	return c.repo.GetByID(ctx, id)
}

func (c *book) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	return c.repo.Update(ctx, book)
}

func (c *book) Delete(ctx context.Context, id uint32) error {
	return c.repo.Delete(ctx, id)
}
