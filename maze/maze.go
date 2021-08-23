package main

import (
	"fmt"
	"os"
)

var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}

	}
	return maze
}
func (p point) add(n point) point {
	return point{p.i + n.i, p.j + n.j}
}
func (p point) at(maze [][]int) (int, bool) {
	val := maze[p.i][p.j]
	if val != 0 {
		return val, false
	}
	return val, true
}

type point struct {
	i, j int
}

func walk(maze [][]int, start, end point) [][]int {
	//create a new var to store steps walked
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))

	}
	// loop an array by appending current and it children
	paths := []point{start}
	for len(paths) > 0 {
		curr := paths[0]
		paths = paths[1:]
		if curr == end {
			break
		}
		// get current node's children by add function
		for _, d := range dirs {
			nxt := curr.add(d)
			// ensure point is not out boundaries
			if nxt.i < 0 || nxt.i > len(maze)-1 || nxt.j < 0 || nxt.j > len(maze[0])-1 {
				continue
			}
			// add function only work when 1.point is not visited 2. point is in map 3. point is not blocked
			// not blocked in maze
			if _, ok := nxt.at(maze); !ok {
				continue
			}
			// not visited in steps
			if _, ok := nxt.at(steps); !ok {
				continue
			}
			if nxt == start {
				continue
			}
			steps[nxt.i][nxt.j] = steps[curr.i][curr.j] + 1

			paths = append(paths, nxt)
		}

	}
	return steps
}

func main() {
	maze := readMaze("maze/maze.txt")
	paths := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range paths {
		for _, val := range row {
			fmt.Printf("%3d", val)

		}
		fmt.Println()
	}
}
