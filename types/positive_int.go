package types

import (
	"fmt"
	"strconv"
)

type Integer interface {
	int | int32 | int64 | int16 | int8
}

type PositiveInt[T Integer] struct {
	value T
}

func NewPositiveInt[T Integer](value T) (PositiveInt[T], error) {
	if value <= 0 {
		return PositiveInt[T]{}, fmt.Errorf("value should be positive: %d", value)
	}

	return PositiveInt[T]{
		value: value,
	}, nil
}

func (p PositiveInt[T]) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, int64(p.value), 10), nil
}

func (p *PositiveInt[T]) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return fmt.Errorf("invalid integer: %w", err)
	}

	if v <= 0 {
		return fmt.Errorf("value should be positive: %d", v)
	}

	p.value = T(v)
	return nil
}
