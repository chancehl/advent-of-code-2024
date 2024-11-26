package utils

import (
	"os"
	"slices"
	"testing"
)

func TestReadInput(t *testing.T) {
	expected := []string{"a", "b", "c", "d"}

	tempDir := os.TempDir()
	tempFile, err := os.CreateTemp(tempDir, "input*.txt")
	if err != nil {
		t.Fatalf("could not create tempfile for test: %v", err)
	}
	if _, err := tempFile.WriteString("a\nb\nc\nd"); err != nil {
		t.Fatalf("could not write to tempfile for test: %v", err)
	}
	defer os.Remove(tempFile.Name())

	actual, err := ReadInput(tempFile.Name())
	if err != nil {
		t.Fatalf("could not read %s: %+v", tempFile.Name(), err)
	}
	tempFile.Close()

	for _, expectedLine := range expected {
		if !slices.Contains(actual, expectedLine) {
			t.Fatalf("line \"%s\" was not present in lines", expectedLine)
		}
	}
}
