package parser

import (
	"encoding/json"
	"fmt"
)

func Parser[T any](body []byte) (*T, error) {
	var data T
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return &data, nil
}
