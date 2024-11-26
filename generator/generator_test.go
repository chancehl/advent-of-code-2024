// Written by ChatGPT because I'm lazy 🤖
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNumberToCamelCase(t *testing.T) {
	tests := []struct {
		input         int
		expectedValue string
		expectedError bool
	}{
		{1, "one", false},
		{10, "ten", false},
		{15, "fifteen", false},
		{25, "twenty_five", false},
		{0, "", true},
		{26, "", true},
	}

	for _, test := range tests {
		result, err := numberToCamelCase(test.input)
		if result != test.expectedValue {
			t.Errorf("numberToCamelCase(%d) = %s; want %s", test.input, result, test.expectedValue)
		}
		if err == nil && test.expectedError {
			t.Errorf("numberToCamelCase(%d) expected to error but did not", test.input)
		}
	}
}

func TestNumberToPascalCase(t *testing.T) {
	tests := []struct {
		input         int
		expectedValue string
		expectedError bool
	}{
		{1, "One", false},
		{10, "Ten", false},
		{15, "Fifteen", false},
		{25, "TwentyFive", false},
		{0, "", true},
		{26, "", true},
	}

	for _, test := range tests {
		result, err := numberToPascalCase(test.input)
		if result != test.expectedValue {
			t.Errorf("numberToPascalCase(%d) = %s; want %s", test.input, result, test.expectedValue)
		}
		if err == nil && test.expectedError {
			t.Errorf("numberToPascalCase(%d) expected to error but did not", test.input)
		}
	}
}

func TestGetAbsolutePath(t *testing.T) {
	relativePath := "testdir/testfile.txt"
	expected := filepath.Join(os.TempDir(), relativePath)

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
