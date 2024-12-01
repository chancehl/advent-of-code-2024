package utils

import (
	"fmt"
)

type SolutionResult struct {
	Result int
	Time   int64
}

func PrintAdventResults(one SolutionResult, two SolutionResult) {
	asciiArt := `ADVENT OF CODE 2024

🎁 Part One 🎁     
🎄 Result: %d 🎄     
🎅 Time:   %d ms 🎅     

🎁 Part Two 🎁     
🎄 Result: %d 🎄     
🎅 Time:   %d ms 🎅     

 🎄🎁 Happy Holidays! 🎁🎄
`
	fmt.Printf(asciiArt, one.Result, one.Time, two.Result, two.Time)
}
