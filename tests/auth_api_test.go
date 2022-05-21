package tests

import (
	"bytes"
	"encoding/json"
	"go-demo/internal/dto"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	var url = `/v1/auth/login`
	var method = `POST`
	tests := []struct {
		name   string
		body   dto.LoginDto
		status int
		want   string
	}{
		{
			name: "Login test",
			body: dto.LoginDto{
				Username: "admin",
				Password: "admin",
			},
			status: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testUrl := url
			bodyStr, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest(method, testUrl, bytes.NewReader(bodyStr))
			res := rest.TestApi(req)
			assert.Equal(t, tt.status, res.Code)
		})
	}
}
