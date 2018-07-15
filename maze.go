package main

// Maze is ...
type Maze struct {
	w      World
	cp     Point
	isGoal bool
}

// NewMaze is ...
func NewMaze() *Maze {
	w := NewWorld()
	cp := Point{x: 0, y: len(w) - 1}
	w[cp.y][cp.x] = CurrentLocation

	return &Maze{w: w, cp: cp, isGoal: false}
}

// Process is ...
func (m *Maze) Process(diff Point) {
	np := m.getNextPoint(diff)
	if np.y < 0 || len(m.w) <= np.y {
		panic("Oops y")
	}
	if np.x < 0 || len(m.w[np.y]) <= np.x {
		panic("Oops x")
	}

	if m.w[np.y][np.x] == Obstacle {
		panic("Oops obstacle")
	}

	m.w[m.cp.y][m.cp.x] = Space
	m.cp = np
	if m.w[np.y][np.x] == Goal {
		m.isGoal = true
		return
	}

	m.w[np.y][np.x] = CurrentLocation
}

// getNextPoint is ...
func (m *Maze) getNextPoint(diff Point) Point {
	nx := m.cp.x + diff.x
	ny := m.cp.y - diff.y

	return Point{x: nx, y: ny}
}
