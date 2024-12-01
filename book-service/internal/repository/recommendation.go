package repository

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"gorm.io/gorm"
)

type RecommendationRepository interface {
	GetRecommendationUserByAuthor(ctx context.Context, userID uint32) ([]model.Book, error)
	GetRecommendationUserByCategory(ctx context.Context, userID uint32) ([]model.Book, error)
	GetRecommendationUserByTitle(ctx context.Context, userID uint32) ([]model.Book, error)
}

type recommendationRepository struct {
	db *gorm.DB
}

func NewRecommendation(db *gorm.DB) RecommendationRepository {
	return &recommendationRepository{db}
}

func (r *recommendationRepository) GetRecommendationUserByAuthor(ctx context.Context, userID uint32) ([]model.Book, error) {
	var borrow []model.Borrowing
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&borrow).Error; err != nil {
		return nil, err
	}

	bookIDs := make([]uint32, 0)
	for _, b := range borrow {
		bookIDs = append(bookIDs, b.BookID)
	}

	var books []model.Book
	if err := r.db.WithContext(ctx).Where("id IN ?", bookIDs).Find(&books).Error; err != nil {
		return nil, err
	}

	authorIDs := make([]uint32, 0)
	for _, b := range books {
		authorIDs = append(authorIDs, b.AuthorID)
	}

	var booksByAuthor []model.Book
	if err := r.db.WithContext(ctx).Where("author_id IN ?", authorIDs).Find(&booksByAuthor).Error; err != nil {
		return nil, err
	}

	return booksByAuthor, nil
}

func (r *recommendationRepository) GetRecommendationUserByCategory(ctx context.Context, userID uint32) ([]model.Book, error) {
	var borrow []model.Borrowing
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&borrow).Error; err != nil {
		return nil, err
	}

	bookIDs := make([]uint32, 0)
	for _, b := range borrow {
		bookIDs = append(bookIDs, b.BookID)
	}

	var books []model.Book
	if err := r.db.WithContext(ctx).Where("id IN ?", bookIDs).Find(&books).Error; err != nil {
		return nil, err
	}

	var categoryBooks []model.CategoryBook
	if err := r.db.WithContext(ctx).Where("book_id IN ?", bookIDs).Find(&categoryBooks).Error; err != nil {
		return nil, err
	}

	categoryIDs := make([]uint32, 0)
	for _, cb := range categoryBooks {
		categoryIDs = append(categoryIDs, cb.CategoryID)
	}

	var booksByCategory []model.Book
	if err := r.db.WithContext(ctx).Where("category_id IN ?", categoryIDs).Find(&booksByCategory).Error; err != nil {
		return nil, err
	}

	return booksByCategory, nil
}

func (r *recommendationRepository) GetRecommendationUserByTitle(ctx context.Context, userID uint32) ([]model.Book, error) {
	var borrow []model.Borrowing
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&borrow).Error; err != nil {
		return nil, err
	}

	bookIDs := make([]uint32, 0)
	for _, b := range borrow {
		bookIDs = append(bookIDs, b.BookID)
	}

	var books []model.Book
	if err := r.db.WithContext(ctx).Where("id IN ?", bookIDs).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
