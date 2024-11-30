package dto

type CategoryBookInput struct {
	CategoryID uint32 `json:"category_id" binding:"required"`
	BookID     uint32 `json:"book_id" binding:"required"`
}
