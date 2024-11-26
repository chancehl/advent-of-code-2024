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

type generatedFiles struct {
	main  string
	test  string
	input string
}

type fileNames generatedFiles
type relativeFilePaths generatedFiles
type absoluteFilePaths generatedFiles

type Case int

const (
	Pascal Case = iota
	Camel
)

func main() {
	day := flag.Int("d", 0, "the day you are completing")

	absoluteTemplatePath := getAbsolutePath("generator/template/day_xxx.txt")
	absoluteTestTemplatePath := getAbsolutePath("generator/template/day_xxx_test.txt")

	template := getTemplateContents(absoluteTemplatePath)
	testTemplate := getTemplateContents(absoluteTestTemplatePath)

	flag.Parse()

	solutionsDirectory := "solutions"

	tokenizedDay := tokenizeDay(*day)
	mainFileContents := tokenizeTemplate(*day, template)
	testFileContents := tokenizeTemplate(*day, testTemplate)

	mainFileName := fmt.Sprintf("%s.go", tokenizedDay)
	testFileName := fmt.Sprintf("%s_test.go", tokenizedDay)
	inputFileName := "input.txt"

	relativePaths := getRelativePaths(solutionsDirectory, tokenizedDay, fileNames{
		main:  mainFileName,
		test:  testFileName,
		input: inputFileName,
	})

	absolutePaths := getAbsoluteFilePaths(relativeFilePaths{
		main:  relativePaths.main,
		test:  relativePaths.test,
		input: relativePaths.input,
	})

	if err := os.MkdirAll(filepath.Join(solutionsDirectory, tokenizedDay), MODE); err != nil {
		log.Fatalf("could not create output directory: %v", err)
	}
	if err := os.WriteFile(absolutePaths.main, []byte(mainFileContents), MODE); err != nil {
		log.Fatalf("could not create output file: %v", err)
	}
	if err := os.WriteFile(absolutePaths.test, []byte(testFileContents), MODE); err != nil {
		log.Fatalf("could not create output test file: %v", err)
	}
	if err := os.WriteFile(absolutePaths.input, []byte{}, MODE); err != nil {
		log.Fatalf("could not create input file: %v", err)
	}

	fmt.Printf(`[success] generated the following files:

- %s/
	- %s/
		- %s
		- %s
		- %s

Happy hacking 🎅`, solutionsDirectory, tokenizedDay, mainFileName, testFileName, inputFileName)
}

func tokenizeDay(day int) string {
	camelCaseDay := formatDay(day, Camel)
	return strings.ReplaceAll("day_xxx", strings.ToLower(TOKEN), camelCaseDay)
}

func tokenizeTemplate(day int, contents string) string {
	pascalCaseDay := formatDay(day, Camel)
	return strings.ReplaceAll(contents, strings.ToUpper(TOKEN), pascalCaseDay)
}

func getRelativePaths(solutionsDir string, dayDir string, fileNames fileNames) (paths relativeFilePaths) {
	paths = relativeFilePaths{
		main:  filepath.Join(solutionsDir, dayDir, fileNames.main),
		test:  filepath.Join(solutionsDir, dayDir, fileNames.test),
		input: filepath.Join(solutionsDir, dayDir, fileNames.input),
	}

	return paths

}

func getAbsoluteFilePaths(relativePaths relativeFilePaths) (paths absoluteFilePaths) {
	paths = absoluteFilePaths{
		main:  getAbsolutePath(relativePaths.main),
		test:  getAbsolutePath(relativePaths.test),
		input: getAbsolutePath(relativePaths.input),
	}

	return paths
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
