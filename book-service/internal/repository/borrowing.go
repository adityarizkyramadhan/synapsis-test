package repository

import (
	"errors"
	"time"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"gorm.io/gorm"
)

type Borrowing interface {
	Borrow(borrow *model.Borrowing) (*model.Borrowing, error)
	Return(borrow *model.Borrowing) (*model.Borrowing, error)
}

type borrowing struct {
	db *gorm.DB
}

func NewBorrowing(db *gorm.DB) *borrowing {
	return &borrowing{db}
}

func (b *borrowing) Borrow(borrow *model.Borrowing) (*model.Borrowing, error) {
	tx := b.db.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	var book model.Book
	if err := tx.Where("id = ?", borrow.BookID).First(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if book.Stock < borrow.Amount {
		tx.Rollback()
		return nil, errors.New("stock is not enough")
	}

	book.Stock -= borrow.Amount
	if err := tx.Save(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	borrow.BorrowedAt = time.Now()

	if err := tx.Create(borrow).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return borrow, nil
}

func (b *borrowing) Return(borrow *model.Borrowing) (*model.Borrowing, error) {
	tx := b.db.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	var book model.Book
	if err := tx.Where("id = ?", borrow.BookID).First(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	book.Stock += borrow.Amount
	if err := tx.Save(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	borrow.ReturnedAt = time.Now()

	if err := tx.Save(borrow).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return borrow, nil
}
