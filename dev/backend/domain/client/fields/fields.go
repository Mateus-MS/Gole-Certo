package fields

// The `Phantom Type` is to the compiler differentiate, preventing the `CPF` to be treated as a `CNPJ`.
type Field[T any] struct {
	value string
}

func (f Field[T]) Get() string {
	return f.value
}

// Generic constructor
func NewField[T any](value string, validate func(string) bool, errInvalid error) (Field[T], error) {
	if !validate(value) {
		return Field[T]{}, errInvalid
	}
	return Field[T]{value: value}, nil
}
