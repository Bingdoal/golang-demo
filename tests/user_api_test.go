package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	var url = `/v1/user`
	var method = `GET`
	var token = loginGetToken()

	tests := []struct {
		name   string
		params map[string]string
		status int
		want   string
	}{
		{
			name: "Get user",
			params: map[string]string{
				"page":     "1",
				"pageSize": "10",
			},
			status: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testUrl := url
			if len(tt.params) > 0 {
				testUrl += "?"
				for k, v := range tt.params {
					testUrl += k + "=" + v + "&"
				}
			}

			req, _ := http.NewRequest(method, testUrl, nil)
			req.Header.Set("Authorization", "Bearer "+token)
			res := rest.TestApi(req)
			assert.Equal(t, tt.status, res.Code)
		})
	}
}
