// Written by ChatGPT because I'm lazy 🤖
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNumberToCamelCase(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "one"},
		{10, "ten"},
		{15, "fifteen"},
		{25, "twenty_five"},
		{0, "invalid_number"},
		{26, "invalid_number"},
	}

	for _, test := range tests {
		result := numberToCamelCase(test.input)
		if result != test.expected {
			t.Errorf("numberToCamelCase(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestNumberToPascalCase(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "One"},
		{10, "Ten"},
		{15, "Fifteen"},
		{25, "TwentyFive"},
		{0, "InvalidNumber"},
		{26, "InvalidNumber"},
	}

	for _, test := range tests {
		result := numberToPascalCase(test.input)
		if result != test.expected {
			t.Errorf("numberToPascalCase(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestGetAbsolutePath(t *testing.T) {
	relativePath := "testdir/testfile.txt"
	expected := filepath.Join(os.TempDir(), relativePath)

	// Create a temporary file to simulate the relative path
	tempDir := os.TempDir()
	testPath := filepath.Join(tempDir, relativePath)
	os.MkdirAll(filepath.Dir(testPath), 0755)
	defer os.RemoveAll(filepath.Dir(testPath))

	absolutePath := getAbsolutePath(testPath)
	if absolutePath != expected {
		t.Errorf("getAbsolutePath(%s) = %s; want %s", testPath, absolutePath, expected)
	}
}

func TestGetTemplateContents(t *testing.T) {
	tempDir := os.TempDir()
	tempFile, err := os.CreateTemp(tempDir, "template*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	expectedContents := "Test template contents"
	if _, err := tempFile.WriteString(expectedContents); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	result := getTemplateContents(tempFile.Name())
	if result != expectedContents {
		t.Errorf("getTemplateContents() = %s; want %s", result, expectedContents)
	}
}
