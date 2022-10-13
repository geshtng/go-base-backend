package dtos

type CreateArticleRequestDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateArticleResponseDTO struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateArticleRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateArticleResponseDTO struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
