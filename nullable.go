package nullable

import (
	"errors"
)

type Primitive interface {
	int | int16 | int32 | int64 | float32 | float64 | string | bool
}

type Nullable[T Primitive] struct {
	content T
	isSet   bool
}

func (v *Nullable[T]) IsNull() bool {
	return !v.isSet
}

func (v *Nullable[T]) Get() T {
	return v.content
}

func (v *Nullable[T]) Set(value T) {
	v.isSet = true
	v.content = value
}

func (v *Nullable[T]) Unset() {
	var zeroValue T

	v.isSet = false
	v.content = zeroValue
}

func (v *Nullable[T]) set_unsafe(value interface{}) error {
	var ok bool
	v.content, ok = value.(T)

	if !ok {
		return errors.New("Invalid value set!")
	}

	return nil
}

func (v *Nullable[T]) asInterface() interface{} {
	return v.content
}

func (Nullable[T]) isNullableStruct() bool {
	return true
}
