package base

type Model interface {
	FindOne() error
	FindAll() ([]Model, error)
	Create() error
	Update() error
	Delete() error
}
