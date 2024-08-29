package dto

type CreateBookDTO struct {
	BookName string `json:"book_name" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Category string `json:"category" binding:"required"`
	Shell    string `json:"shell" binding:"required"`
}

type UpdateBookDTO struct {
	BookName *string `json:"book_name,omitempty"`
	Author   *string `json:"author,omitempty"`
	Category *string `json:"category,omitempty"`
	Shell    *string `json:"shell,omitempty"`
}
