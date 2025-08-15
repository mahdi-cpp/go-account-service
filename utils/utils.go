package utils

import (
	"encoding/json"
	"fmt"
)

func ToStringJson[T any](data T) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal struct to JSON: %w", err)
	}
	return string(jsonBytes), nil
}
