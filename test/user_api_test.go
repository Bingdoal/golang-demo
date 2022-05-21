package tests

import "testing"

func TestUserApi(t *testing.T) {
	type api struct {
		url           string
		method        string
		header        map[string]string
		body          string
		exceptStatus  int
		exceptMessage string
	}
	tests := []struct {
		name string
		apis api
		want string
	}{
		{
			name: "Test Case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
