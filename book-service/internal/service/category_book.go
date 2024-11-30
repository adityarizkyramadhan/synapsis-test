package service

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/repository"
)

type CategoryBookService interface {
	Add(ctx context.Context, category *model.CategoryBook) (*model.CategoryBook, error)
	Delete(ctx context.Context, id uint32) error
}

type categoryBook struct {
	repo repository.CategoryBookRepository
}

func NewCategoryBook(repo repository.CategoryBookRepository) CategoryBookService {
	return &categoryBook{repo}
}

func (c *categoryBook) Add(ctx context.Context, category *model.CategoryBook) (*model.CategoryBook, error) {
	return c.repo.Add(ctx, category)
}

func (c *categoryBook) Delete(ctx context.Context, id uint32) error {
	return c.repo.Delete(ctx, id)
}

func (c *categoryBook) GetAll(ctx context.Context) ([]model.CategoryBook, error) {
	return c.repo.GetAll(ctx)
}
