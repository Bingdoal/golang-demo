package dto

type PostDto struct {
	Content  string `json:"content" validate:"required,min=2"`
	AuthorID uint   `json:"authorId" validate:"required"`
}

