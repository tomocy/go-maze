package main

// Up is ...
// Left is ...
// Right is ...
// Down is ...
const (
	Up    = "j"
	Left  = "h"
	Right = "l"
	Down  = "k"
)

// MazeController is ...
type MazeController struct {
	cursor map[string]Point
	m      *Maze
	mv     *MazeView
}

// NewMazeController is ...
func NewMazeController() MazeController {
	cursor := map[string]Point{
		Up:    Point{x: 0, y: 1},
		Left:  Point{x: -1, y: 0},
		Right: Point{x: 1, y: 0},
		Down:  Point{x: 0, y: -1},
	}
	m := NewMaze()
	mv := NewMazeView()

	return MazeController{cursor: cursor, m: m, mv: mv}
}

// Process is ...
func (mc *MazeController) Process(cmd string) string {
	p, ok := mc.cursor[cmd]
	if !ok {
		panic("Oops cmd")
	}

	mc.m.Process(p)

	return mc.mv.Render(mc.m)
}
