package dto

type CategoryInput struct {
	Name string `json:"name" binding:"required,min=3,max=255"`
}

type CategoryOutput struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
