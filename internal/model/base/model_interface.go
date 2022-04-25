package base

type Models interface {
	FindAll() error
}

type Model interface {
	FindOne() error
	Create() error
	Update() error
	Delete() error
}
