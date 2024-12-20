package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const MODE = 0777

type Case int

const (
	Pascal Case = iota
	Camel
)

type flags struct {
	day int
}

func main() {
	flags := parseFlags()

	day := flags.day

	absoluteTemplatePath := getAbsolutePath("generator/template/main.go")
	absoluteTestTemplatePath := getAbsolutePath("generator/template/main_test.go")

	template := getTemplateContents(absoluteTemplatePath)
	testTemplate := getTemplateContents(absoluteTestTemplatePath)

	solutionsDir := "solutions"

	tokenizedDay := tokenizeDay(day)
	mainFileContents := tokenizeTemplate(day, template)
	testFileContents := tokenizeTemplate(day, testTemplate)

	mainFileName := fmt.Sprintf("%s.go", tokenizedDay)
	testFileName := fmt.Sprintf("%s_test.go", tokenizedDay)
	inputFileName := "input.txt"

	outputDir := filepath.Join(solutionsDir, tokenizedDay)

	mainFile := getAbsolutePath(filepath.Join(outputDir, mainFileName))
	testFile := getAbsolutePath(filepath.Join(outputDir, testFileName))
	inputFile := getAbsolutePath(filepath.Join(outputDir, inputFileName))

	if err := os.MkdirAll(filepath.Join(outputDir), MODE); err != nil {
		log.Fatalf("could not create output directory: %v", err)
	}
	if err := os.WriteFile(mainFile, []byte(mainFileContents), MODE); err != nil {
		log.Fatalf("could not create main file: %v", err)
	}
	if err := os.WriteFile(testFile, []byte(testFileContents), MODE); err != nil {
		log.Fatalf("could not create test file: %v", err)
	}
	if err := os.WriteFile(inputFile, []byte{}, MODE); err != nil {
		log.Fatalf("could not create input file: %v", err)
	}

	fmt.Printf(`[success] generated the following files:

- %s/
	- %s/
		- %s
		- %s
		- %s

Happy hacking 🎅`, solutionsDir, tokenizedDay, mainFileName, testFileName, inputFileName)
}

func parseFlags() flags {
	day := flag.Int("d", 0, "the day you are completing")
	flag.Parse()

	return flags{day: *day}
}

func tokenizeDay(day int) string {
	camelCaseDay := formatDay(day, Camel)
	return strings.ReplaceAll("day_xxx", "xxx", camelCaseDay)
}

func tokenizeTemplate(day int, contents string) string {
	pascalCaseDay := formatDay(day, Pascal)
	camelCaseDay := formatDay(day, Camel)

	pascalized := strings.ReplaceAll(contents, "XXX", pascalCaseDay)
	camelized := strings.ReplaceAll(pascalized, "xxx", camelCaseDay)

	return camelized
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

func numberToCamelCase(num int) (string, error) {
	if num < 1 || num > 25 {
		return "", fmt.Errorf("invalid input: %d", num)
	}

	values := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty_one", "twenty_two", "twenty_three", "twenty_four", "twenty_five"}

	return values[num-1], nil
}

func numberToPascalCase(num int) (string, error) {
	if num < 1 || num > 25 {
		return "", fmt.Errorf("invalid input: %d", num)
	}

	values := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen", "Twenty", "TwentyOne", "TwentyTwo", "TwentyThree", "TwentyFour", "TwentyFive"}

	return values[num-1], nil
}

func formatDay(num int, c Case) string {
	if c == Camel {
		camelCaseDay, err := numberToCamelCase(num)
		if err != nil {
			log.Fatalf("could not format date: %v", err)
		}
		return camelCaseDay
	}

	pascalCaseDay, err := numberToPascalCase(num)
	if err != nil {
		log.Fatalf("could not format date: %v", err)
	}
	return pascalCaseDay
}
