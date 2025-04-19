package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

const DefaultPageSize = 20

// EncodeCursor encodes any struct into a base64 URL-safe string.
func EncodeCursor[T any](cursor T) (string, error) {
	jsonBytes, err := json.Marshal(cursor)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(jsonBytes), nil
}

// DecodeCursor decodes a base64 URL-safe string into the provided struct type.
func DecodeCursor[T any](encoded string) (*T, error) {
	if encoded == "" {
		return nil, errors.New("empty cursor")
	}
	data, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	var cursor T
	if err := json.Unmarshal(data, &cursor); err != nil {
		return nil, err
	}
	return &cursor, nil
}
