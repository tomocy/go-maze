package main

// CurrentLocation is ...
// Space is ...
// Obstacle is ...
const (
	CurrentLocation    level = -1
	CurrentLocationStr       = "@"
	Space              level = 0
	SpaceStr                 = "."
	Obstacle           level = 1
	ObstacleStr              = "#"
	Goal               level = 2
	GoalStr                  = "G"
)

// Level is ...
type Level interface {
	String() string
}

type level int

func (l level) String() string {
	switch l {
	case CurrentLocation:
		return CurrentLocationStr
	case Space:
		return SpaceStr
	case Obstacle:
		return ObstacleStr
	case Goal:
		return GoalStr
	default:
		panic("Oops level")
	}
}
