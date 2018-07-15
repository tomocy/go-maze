package main

import (
	"bufio"
	"os"
)

func main() {
	mc := NewMazeController()
	scanner := bufio.NewScanner(os.Stdin)

	print(mc.mv.Render(mc.m))
	for scanner.Scan() {
		cmd := scanner.Text()
		print(mc.Process(cmd))
		if mc.m.isGoal {
			break
		}
	}
}
