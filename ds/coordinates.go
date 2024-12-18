package ds

import "strings"

type Coordinates struct {
	Row int
	Col int
}

func NewCoordinate(row, col int) Coordinates {
	return Coordinates{Row: row, Col: col}
}

// Implement the LessThan method for Coordinates
func (c Coordinates) LessThan(other Coordinates) bool {
	if c.Row < other.Row {
		return true
	}
	if c.Row == other.Row && c.Col < other.Col {
		return true
	}
	return false
}

func CoordinateComparator(a, b Coordinates) bool {
	return a.LessThan(b)
}

func (c Coordinates) Equals(other Coordinates) bool {
	return c.Row == other.Row && c.Col == other.Col
}

// Returns the coordinate "nieghbors" (up, down, left, right)
//
// Note: this method makes no gaurantees that the coordinates returned
// are within the bounds of the matrix
func GetNeighbors(c Coordinates) []Coordinates {
	return []Coordinates{
		{Row: c.Row - 1, Col: c.Col}, // up
		{Row: c.Row + 1, Col: c.Col}, // down
		{Row: c.Row, Col: c.Col - 1}, // left
		{Row: c.Row, Col: c.Col + 1}, // right
	}
}

// Returns the coordinate "nieghbors" (up, down, left, right) if they are in bounds
func GetInBoundsNeighbors[T int | string](c Coordinates, matrix Matrix[T]) []Coordinates {
	valid := []Coordinates{}
	for _, neighbcr := range GetNeighbors(c) {
		if matrix.IsCoordInBounds(neighbcr) {
			valid = append(valid, neighbcr)
		}
	}
	return valid
}

func GetNeighborsInOrder(c Coordinates, order string) []Coordinates {
	up := Coordinates{Row: c.Row - 1, Col: c.Col}    // up
	down := Coordinates{Row: c.Row + 1, Col: c.Col}  // down
	left := Coordinates{Row: c.Row, Col: c.Col - 1}  // left
	right := Coordinates{Row: c.Row, Col: c.Col + 1} // right

	ordered := []Coordinates{}

	for _, char := range strings.Split(order, "") {
		if char == "U" {
			ordered = append(ordered, up)
		} else if char == "D" {
			ordered = append(ordered, down)
		} else if char == "L" {
			ordered = append(ordered, left)
		} else if char == "R" {
			ordered = append(ordered, right)
		} else {
			panic("character needs to be one of: U, D, L, R")
		}
	}

	return ordered
}
