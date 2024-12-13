package ds

type Coordinates struct {
	Row int
	Col int
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
