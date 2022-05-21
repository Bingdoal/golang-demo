package common

import "fmt"

type StatusError struct {
	Status  int
	Message string
}

func (err StatusError) Error() string {
	return fmt.Sprintf("%d: %s", err.Status, err.Message)
}
