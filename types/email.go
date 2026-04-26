package types

import (
	"encoding/json"
	"errors"
	"net/mail"
	"strings"
)

type Email struct {
	value string
}

func (e Email) Value() string {
	return e.value
}

func NewEmail(raw string) (Email, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return Email{}, errors.New("email cannot be empty")
	}

	// Нормализуем
	raw = strings.ToLower(raw)

	parsed, err := mail.ParseAddress(raw)
	if err != nil {
		return Email{}, errors.New("invalid email format")
	}

	if parsed.Address != raw {
		return Email{}, errors.New("invalid email format")
	}

	return Email{
		value: raw,
	}, nil
}

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.value)
}

func (e *Email) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := NewEmail(s)
	if err != nil {
		return err
	}

	*e = parsed
	return nil
}
