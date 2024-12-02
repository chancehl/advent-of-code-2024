package random

import "math/rand"

func Choice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}
