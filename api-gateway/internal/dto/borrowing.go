package dto

type Borrowing struct {
	BookID uint32 `json:"book_id" binding:"required"`
	Amount uint32 `json:"amount" binding:"required,min=1"`
}
