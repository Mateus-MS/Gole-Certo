package user

type User interface {
	GetIdentifier() string
	IsValid() error
}
