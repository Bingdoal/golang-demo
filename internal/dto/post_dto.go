package dto

type PostDto struct {
	Content  string `json:"content"`
	AuthorID uint   `json:"authorId"`
}
