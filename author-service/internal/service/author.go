package service

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/author-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/author-service/internal/repository"
)

type Author struct {
	repoAuthor repository.AuthorRepository
}

type AuthorService interface {
	Create(ctx context.Context, author *model.Author) error
	ReadAll(ctx context.Context) ([]model.Author, error)
	ReadByID(ctx context.Context, id int) (model.Author, error)
	Update(ctx context.Context, author *model.Author) error
	Delete(ctx context.Context, id int) error
}

func NewAuthor(repoAuthor repository.AuthorRepository) AuthorService {
	return &Author{repoAuthor: repoAuthor}
}

func (a *Author) Create(ctx context.Context, author *model.Author) error {
	return a.repoAuthor.Create(ctx, author)
}

func (a *Author) ReadAll(ctx context.Context) ([]model.Author, error) {
	return a.repoAuthor.ReadAll(ctx)
}

func (a *Author) ReadByID(ctx context.Context, id int) (model.Author, error) {
	return a.repoAuthor.ReadByID(ctx, id)
}

func (a *Author) Update(ctx context.Context, author *model.Author) error {
	return a.repoAuthor.Update(ctx, author)
}

func (a *Author) Delete(ctx context.Context, id int) error {
	return a.repoAuthor.Delete(ctx, id)
}
