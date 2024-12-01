package dto

type BookInput struct {
	Title       string `json:"title" binding:"required,min=3,max=255"`
	AuthorID    uint32 `json:"author_id" binding:"required"`
	Description string `json:"description" binding:"required,min=3"`
	Year        uint32 `json:"year" binding:"required"`
	Stock       uint32 `json:"stock" binding:"required"`
}

type BookOutput struct {
	ID          uint32 `json:"id"`
	Title       string `json:"title"`
	AuthorID    uint32 `json:"author_id"`
	Description string `json:"description"`
	Year        uint32 `json:"year"`
	Stock       uint32 `json:"stock"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
