package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}
	return strings.Split(string(data), "\n"), nil
}
