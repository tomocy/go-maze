package main

// MazeView is ...
type MazeView struct {
}

// NewMazeView is ...
func NewMazeView() *MazeView {
	return &MazeView{}
}

// Render is ...
func (mv *MazeView) Render(m *Maze) string {
	if m.isGoal {
		return "Goal\n"
	}

	world := mv.renderWorld(m.w)
	return string(world)
}

func (mv *MazeView) renderWorld(w World) []byte {
	b := make([]byte, 0, 10)
	for _, y := range w {
		for _, x := range y {
			b = append(b, x.String()...)
		}
		b = append(b, "\n"...)
	}

	return b
}
