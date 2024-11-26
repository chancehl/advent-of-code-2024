package utils

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func ReadFileFromFs(fsys fs.FS, path string) ([]string, error) {
	bytes, err := fs.ReadFile(fsys, path)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(bytes), "\n"), nil
}

func ReadInputFile(path string) ([]string, error) {
	realFS := os.DirFS(".")
	return ReadFileFromFs(realFS, path)
}

func ReadFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}
	return strings.Split(string(data), "\n"), nil
}
