package types

import (
	"encoding/json"
	"errors"
	"strings"
)

type NonEmptyString struct {
	value string
}

func (n NonEmptyString) Value() string {
	return n.value
}

func NewNonEmptyString(s string) (NonEmptyString, error) {
	trimmed := strings.TrimSpace(s)

	if len(trimmed) == 0 {
		return NonEmptyString{}, errors.New("string shouldn't be empty")
	}

	return NonEmptyString{
		value: trimmed,
	}, nil
}

func (n *NonEmptyString) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	validated, err := NewNonEmptyString(s)
	if err != nil {
		return err
	}

	*n = validated
	return nil
}

func (n *NonEmptyString) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.value)
}
