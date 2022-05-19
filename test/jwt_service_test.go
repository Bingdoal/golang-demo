package tests

import (
	"go-demo/internal/service/jwt_service"
	"reflect"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		subject string
		body    map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				subject: "userId",
				body: map[string]string{
					"hello": "world",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			token := jwt_service.GenerateToken(tt.args.subject, tt.args.body)
			subject, body, err := jwt_service.ValidateToken(token)
			if err != nil {
				t.Errorf("Generate token is invalid. err: %s", err.Error())
			}
			if subject != tt.args.subject {
				t.Errorf("Resolve subject is invalid. return: %s, input: %s",
					subject, tt.args.subject)
			}
			if !reflect.DeepEqual(body, tt.args.body) {
				t.Errorf("Resolve body is invalid. return: %s, input: %s",
					body, tt.args.body)
			}
		})
	}
}
