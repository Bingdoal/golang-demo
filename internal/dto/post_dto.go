package dto

type PostDto struct {
	Content  string `json:"content" binding:"required,min=2"`
	AuthorID uint64 `json:"authorId" binding:"required"`
}
