package dto

type PostDto struct {
	Content  string `json:"content"`
	AuthorID uint64 `json:"authorId"`
}
