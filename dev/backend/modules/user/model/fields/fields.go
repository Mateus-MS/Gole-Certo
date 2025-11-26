package fields

import (
	"encoding/json"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

var (
	ErrInvalidAddress = errors.New("invalid address")
)

type Validator interface {
	Validate(value string) bool
}

// The `Phantom Type` is to the compiler differentiate, preventing the `CPF` to be treated as a `CNPJ`.
type Field[T Validator] struct {
	value string
	Tag   T
}

func (f *Field[T]) Get() string {
	return f.value
}

// Generic constructor
func NewField[T Validator](value string, validate func(string) bool, errInvalid error) (Field[T], error) {
	if !validate(value) {
		return Field[T]{}, errInvalid
	}
	return Field[T]{value: value}, nil
}

// Override the default json's methods for marshan and unmarshal
func (f *Field[T]) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return errors.New("invalid value for field: " + err.Error())
	}

	if !f.Tag.Validate(raw) {
		return fmt.Errorf("invalid value for %T: %s", f.Tag, raw)
	}

	f.value = raw
	return nil
}

func (f Field[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.value)
}

// Override the default bson's methods to use with mongoDB
func (f *Field[T]) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	var raw string
	if err := bson.UnmarshalValue(t, data, &raw); err != nil {
		return err
	}

	if !f.Tag.Validate(raw) {
		return fmt.Errorf("invalid value for %T: %s", f.Tag, raw)
	}

	f.value = raw
	return nil
}

func (f Field[T]) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(f.value)
}
