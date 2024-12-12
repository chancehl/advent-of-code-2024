package ds

import "fmt"

type Coordinates struct {
	Row int
	Col int
}

func (coords *Coordinates) Key() string {
	return fmt.Sprintf("%d,%d", coords.Row, coords.Col)
}
