package utils

import (
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
