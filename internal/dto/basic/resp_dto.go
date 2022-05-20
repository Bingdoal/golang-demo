package basic

import "go-demo/internal/enum"

type RespDto struct {
	Message    enum.MessageType `json:"message"`
	Data       interface{}      `json:"data,omitempty"`
	Err        string           `json:"err,omitempty"`
	Pagination *Pagination       `json:"pagenation,omitempty"`
}
