package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	i, j int
}

type mazeSize struct {
	x, y int
}

func readMaze(filepath string) (map[point]int, mazeSize) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	var maze = make(map[point]int)
	var size mazeSize
	lines := strings.TrimSpace(string(file))
	for i, line := range strings.Split(lines, "\n") {
		var p point
		p.i = i
		for j, v := range strings.Split(line, " ") {
			p.j = j
			maze[p], _ = strconv.Atoi(v)
			size.x = j
			fmt.Printf("[%d %d %d]", i, j, maze[p])
		}
		size.y = i
		fmt.Println()
	}
	return maze, size
}

func (p point) moveDirection(d point, maxi, maxj int) (point, error) {
	var ret point
	ret.i, ret.j = p.i+d.i, p.j+d.j
	if ret.i < 0 || ret.j < 0 || ret.i > maxi-1 || ret.j > maxj-1 {
		return ret, errors.New("Bad direction")
	}
	return ret, nil
}

var directions []point = []point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func walk(maze map[point]int, start point, size mazeSize) []point {
	var steps []point
	nexts := []point{start}
	walked := make(map[point]bool)
	for len(nexts) > 0 {
		cur := nexts[0]
		nexts = nexts[1:]
		for _, direction := range directions {
			next, err := cur.moveDirection(direction, size.x, size.y)
			if err != nil || maze[next] > 0 || walked[next] {
				continue
			}
			nexts = append(nexts, next)
		}
		steps = append(steps, cur)
		walked[cur] = true
	}
	return steps
}

func main() {
	maze, size := readMaze("input1.txt")
	fmt.Println(walk(maze, point{0, 0}, size))
}
