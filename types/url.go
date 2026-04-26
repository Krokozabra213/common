package types

import (
	"encoding/json"
	"errors"
	"net/url"
	"path"
	"strings"
)

type URL struct {
	value string
}

func (u URL) Value() string {
	return u.value
}

func NewURL(raw string) (URL, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return URL{}, errors.New("url is empty")
	}

	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		return URL{}, errors.New("invalid url format")
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return URL{}, errors.New("invalid url scheme: only http and https are allowed")
	}

	if parsed.Host == "" {
		return URL{}, errors.New("invalid url: missing host")
	}

	if parsed.Path != "" {
		parsed.Path = path.Clean(parsed.Path)
	}

	return URL{
		value: parsed.String(),
	}, nil
}

func (u URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

func (u *URL) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := NewURL(s)
	if err != nil {
		return err
	}

	*u = parsed
	return nil
}
