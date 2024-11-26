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

	camelCaseDay := formatDay(*day, Camel)
	pascalCaseDay := formatDay(*day, Pascal)

	solutionsDirectory := "solutions"

	dayDirectory := strings.ReplaceAll("day_xxx", strings.ToLower(TOKEN), camelCaseDay)
	fileName := strings.ReplaceAll("day_xxx", strings.ToLower(TOKEN), camelCaseDay)
	mainFileContents := strings.ReplaceAll(template, strings.ToUpper(TOKEN), pascalCaseDay)
	testFileContents := strings.ReplaceAll(testTemplate, strings.ToUpper(TOKEN), pascalCaseDay)

	mainFileName := fmt.Sprintf("%s.go", fileName)
	testFileName := fmt.Sprintf("%s_test.go", fileName)
	inputFileName := "input.txt"

	relativeMainFilePath := filepath.Join(solutionsDirectory, dayDirectory, mainFileName)
	relativeTestFilePath := filepath.Join(solutionsDirectory, dayDirectory, testFileName)
	relativeInputFilePath := filepath.Join(solutionsDirectory, dayDirectory, inputFileName)

	absoluteMainFile := getAbsolutePath(relativeMainFilePath)
	absoluteTestFile := getAbsolutePath(relativeTestFilePath)
	absoluteInputFile := getAbsolutePath(relativeInputFilePath)

	if err := os.MkdirAll(filepath.Join(solutionsDirectory, dayDirectory), MODE); err != nil {
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

	fmt.Printf(`[success] generated the following files:

- %s/
	- %s/
		- %s
		- %s
		- %s

Happy hacking 🎅`, solutionsDirectory, dayDirectory, mainFileName, testFileName, inputFileName)
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
