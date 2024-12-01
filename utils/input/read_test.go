package input

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	expected := "a\nb\nc\nd"

	tempDir := os.TempDir()
	tempFile, err := os.CreateTemp(tempDir, "input*.txt")
	if err != nil {
		t.Fatalf("could not create tempfile for test: %v", err)
	}
	if _, err := tempFile.WriteString("a\nb\nc\nd"); err != nil {
		t.Fatalf("could not write to tempfile for test: %v", err)
	}
	defer os.Remove(tempFile.Name())

	actual, err := Read(tempFile.Name())
	if err != nil {
		t.Fatalf("could not read %s: %+v", tempFile.Name(), err)
	}
	tempFile.Close()

	if actual != expected {
		t.Fatalf("actual \"%s\" did not match expected \"%s\"", actual, expected)
	}
}
