package dto

type AuthorInput struct {
	Name  string `json:"name" binding:"required,min=3,max=100"`
	Email string `json:"email" binding:"required,email"`
	Bio   string `json:"bio" binding:"required,min=10,max=250"`
}

type AuthorOutput struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
