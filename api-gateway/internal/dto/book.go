package dto

type BookInput struct {
	Title       string `json:"title" binding:"required,min=3,max=255"`
	AuthorID    uint32 `json:"author_id" binding:"required"`
	Description string `json:"description" binding:"required,min=3"`
	Year        uint32 `json:"year" binding:"required"`
	Stock       uint32 `json:"stock" binding:"required"`
}
