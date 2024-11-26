package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const TOKEN = "xxx"
const MODE = 0777

func main() {
	day := flag.Int("d", 0, "the day you are completing")

	absoluteTemplatePath := getAbsolutePath("generator/template/day_xxx.txt")
	absoluteTestTemplatePath := getAbsolutePath("generator/template/day_xxx_test.txt")

	template := getTemplateContents(absoluteTemplatePath)
	testTemplate := getTemplateContents(absoluteTestTemplatePath)

	flag.Parse()

	directoryName := strings.ReplaceAll("day_xxx", strings.ToLower(TOKEN), numberToCamelCase(*day))
	fileName := strings.ReplaceAll("day_xxx", strings.ToLower(TOKEN), numberToCamelCase(*day))
	mainFileContents := strings.ReplaceAll(template, strings.ToUpper(TOKEN), numberToPascalCase(*day))
	testFileContents := strings.ReplaceAll(testTemplate, strings.ToUpper(TOKEN), numberToPascalCase(*day))

	mainFileName := fmt.Sprintf("%s.go", fileName)
	testFileName := fmt.Sprintf("%s_test.go", fileName)
	inputFileName := "input.txt"

	relativeMainFilePath := filepath.Join(directoryName, mainFileName)
	relativeTestFilePath := filepath.Join(directoryName, testFileName)
	relativeInputFilePath := filepath.Join(directoryName, inputFileName)

	absoluteMainFile := getAbsolutePath(relativeMainFilePath)
	absoluteTestFile := getAbsolutePath(relativeTestFilePath)
	absoluteInputFile := getAbsolutePath(relativeInputFilePath)

	if err := os.Mkdir(directoryName, MODE); err != nil {
		log.Fatalf("could not create output directory: %v", err)
	}
	if err := os.WriteFile(absoluteMainFile, []byte(mainFileContents), MODE); err != nil {
		log.Fatalf("could not create output file: %v", err)
	}
	if err := os.WriteFile(absoluteTestFile, []byte(testFileContents), MODE); err != nil {
		log.Fatalf("could not create output test file: %v", err)
	}
	if err := os.WriteFile(absoluteInputFile, []byte{}, MODE); err != nil {
		log.Fatalf("could not create input file: %v", err)
	}

	fmt.Printf(`[success] Generated the following files:

- %s/
	- %s
	- %s
	- %s

Happy hacking 🎅`, directoryName, mainFileName, testFileName, inputFileName)
}

func getAbsolutePath(relative string) string {
	absolutePath, err := filepath.Abs(relative)
	if err != nil {
		log.Fatalf("error constructing absolute path: %v", err)
	}
	return absolutePath
}

func getTemplateContents(path string) string {
	template, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("could not read template file: %v", err)
	}
	return string(template)
}

func numberToCamelCase(num int) string {
	if num < 1 || num > 25 {
		return "invalid_number"
	}

	// Define the base number words
	units := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	teens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "", "twenty"}

	var words []string

	if num >= 10 && num < 20 {
		// Handle teens
		words = append(words, teens[num-10])
	} else {
		// Handle tens and units
		if num >= 20 {
			words = append(words, tens[num/10])
		}
		if num%10 > 0 {
			words = append(words, units[num%10])
		}
	}

	// Join words with underscores
	return strings.Join(words, "_")
}

func numberToPascalCase(num int) string {
	if num < 1 || num > 25 {
		return "InvalidNumber"
	}

	// Define the base number words
	units := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teens := []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "", "Twenty"}

	var words []string

	if num >= 10 && num < 20 {
		// Handle teens
		words = append(words, teens[num-10])
	} else {
		// Handle tens and units
		if num >= 20 {
			words = append(words, tens[num/10])
		}
		if num%10 > 0 {
			words = append(words, units[num%10])
		}
	}

	// Join words without separators to form PascalCase
	return strings.Join(words, "")
}
