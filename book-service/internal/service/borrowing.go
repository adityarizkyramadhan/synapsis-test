package service

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/repository"
)

type Borrowing interface {
	Borrow(ctx context.Context, borrow *model.Borrowing) (*model.Borrowing, error)
	Return(ctx context.Context, borrow *model.Borrowing) (*model.Borrowing, error)
}

type borrowing struct {
	repo repository.Borrowing
}

func NewBorrowing(repo repository.Borrowing) *borrowing {
	return &borrowing{repo}
}

func (b *borrowing) Borrow(ctx context.Context, borrow *model.Borrowing) (*model.Borrowing, error) {
	return b.repo.Borrow(ctx, borrow)
}

func (b *borrowing) Return(ctx context.Context, borrow *model.Borrowing) (*model.Borrowing, error) {
	return b.repo.Return(ctx, borrow)
}
