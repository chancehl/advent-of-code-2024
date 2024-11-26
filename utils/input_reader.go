package utils

import (
	"fmt"
	"os"
)

func ReadInput(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not read file: %v", err)
	}
	return string(data), nil
}
