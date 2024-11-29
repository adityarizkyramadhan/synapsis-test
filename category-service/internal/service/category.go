package service

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/category-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/category-service/internal/repository"
)

type Category struct {
	repoCategory repository.CategoryRepository
}

type CategoryService interface {
	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id uint32) error
	GetByID(ctx context.Context, id uint32) (*model.Category, error)
	GetAll(ctx context.Context) ([]model.Category, error)
}

func NewCategory(repoCategory repository.CategoryRepository) CategoryService {
	return &Category{repoCategory}
}

func (c *Category) Create(ctx context.Context, category *model.Category) error {
	return c.repoCategory.Create(ctx, category)
}

func (c *Category) Update(ctx context.Context, category *model.Category) error {
	return c.repoCategory.Update(ctx, category)
}

func (c *Category) Delete(ctx context.Context, id uint32) error {
	return c.repoCategory.Delete(ctx, id)
}

func (c *Category) GetByID(ctx context.Context, id uint32) (*model.Category, error) {
	return c.repoCategory.GetByID(ctx, id)
}

func (c *Category) GetAll(ctx context.Context) ([]model.Category, error) {
	return c.repoCategory.GetAll(ctx)
}
