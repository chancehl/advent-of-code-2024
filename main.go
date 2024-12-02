package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/random"
)

func main() {
	days, err := filepath.Glob(filepath.Join("solutions", "day*"))
	if err != nil {
		fmt.Printf("error finding solutions: %+v", err)
		return
	}

	for _, day := range days {
		baseName := filepath.Base(day)
		goFilePath := filepath.Join(day, fmt.Sprintf("%s.go", strings.ToLower(baseName)))

		emoji := random.Choice([]string{"🎄", "🎅", "🎁", "🧝"})
		fmt.Printf("%s running %s...\n", emoji, goFilePath)

		cmd := exec.Command("go", "run", goFilePath)
		output, _ := cmd.CombinedOutput()

		fmt.Println(string(output))
	}
}
