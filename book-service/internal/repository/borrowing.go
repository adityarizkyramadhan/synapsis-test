package repository

import (
	"context"
	"errors"
	"time"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"gorm.io/gorm"
)

type Borrowing interface {
	Borrow(ctx context.Context, borrow *model.Borrowing) (*model.Borrowing, error)
	Return(ctx context.Context, borrow *model.Borrowing) (*model.Borrowing, error)
}

type borrowing struct {
	db *gorm.DB
}

func NewBorrowing(db *gorm.DB) *borrowing {
	return &borrowing{db}
}

func (b *borrowing) Borrow(ctx context.Context, borrow *model.Borrowing) (*model.Borrowing, error) {
	tx := b.db.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	var book model.Book
	if err := tx.WithContext(ctx).Where("id = ?", borrow.BookID).First(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if book.Stock < borrow.Amount {
		tx.Rollback()
		return nil, errors.New("stock is not enough")
	}

	book.Stock -= borrow.Amount
	if err := tx.WithContext(ctx).Save(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	borrow.BorrowedAt = time.Now()

	if err := tx.WithContext(ctx).Create(borrow).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return borrow, nil
}

func (b *borrowing) Return(ctx context.Context, borrow *model.Borrowing) (*model.Borrowing, error) {
	tx := b.db.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	var borrowData model.Borrowing
	if err := tx.WithContext(ctx).Where("id = ?", borrow.ID).First(&borrowData).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var book model.Book
	if err := tx.WithContext(ctx).Where("id = ?", borrowData.BookID).First(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	book.Stock += borrowData.Amount
	if err := tx.WithContext(ctx).Save(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	borrowData.ReturnedAt = time.Now()

	if err := tx.WithContext(ctx).Save(&borrowData).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &borrowData, nil
}
