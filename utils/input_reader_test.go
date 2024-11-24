package utils

import (
	"slices"
	"testing"
	"testing/fstest"
)

func TestInputReader(t *testing.T) {
	path := "test_input.txt"
	expected := []string{"a", "b", "c", "d"}

	mockFS := fstest.MapFS{
		path: {
			Data: []byte("a\nb\nc\nd"),
		},
	}

	inputFileLines, err := ReadFileFromFs(mockFS, path)
	if err != nil {
		t.Fatalf("could not read test_input.txt: %+v", err)
	}

	for _, expectedLine := range expected {
		if !slices.Contains(inputFileLines, expectedLine) {
			t.Fatalf("line \"%s\" was not present in lines", expectedLine)
		}
	}
}
