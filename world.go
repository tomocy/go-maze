package main

// World is ...
type World [][]Level

// NewWorld is ...
func NewWorld() World {
	return [][]Level{
		[]Level{Obstacle, Goal, Obstacle, Obstacle, Obstacle},
		[]Level{Obstacle, Space, Obstacle, Space, Space},
		[]Level{Obstacle, Space, Obstacle, Space, Obstacle},
		[]Level{Obstacle, Space, Space, Space, Obstacle},
		[]Level{Space, Space, Obstacle, Obstacle, Obstacle},
	}
}
