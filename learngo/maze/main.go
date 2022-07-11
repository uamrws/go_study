package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) (maze [][]int, steps [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, err = fmt.Fscan(file, &row, &col)
	if err != nil {
		panic(err)
	}
	maze = make([][]int, row)
	steps = make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		steps[i] = make([]int, col)
		for j := range maze[i] {
			_, err = fmt.Fscan(file, &maze[i][j])
			if err != nil {
				panic(err)
			}
		}
	}
	return
}

type point struct {
	i, j int
}

func (p *point) add(o *point) point {
	return point{p.i + o.i, p.j + o.j}
}

func (p *point) at(dest [][]int) (int, bool) {
	if p.i >= len(dest) || p.i < 0 || p.j >= len(dest) || p.j < 0 {
		return 0, false
	}
	return dest[p.i][p.j], true
}

var directions = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func walk(maze, steps [][]int) {
	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		// 每个点朝上下左右四个方向延伸一格
		for _, dir := range directions {
			next := cur.add(&dir)
			// 查看新格子的情况
			// 是否能通过（在maze中为0就是通过）
			if val, ok := next.at(maze); val != 0 || !ok {
				continue
			}
			// 是否已经延伸过（在steps中不为0）
			if val, ok := next.at(steps); val != 0 || !ok {
				continue
			}
			// 也不能回到起点
			if next == start {
				continue
			}

			// 记录下到这个点用的step数
			steps[next.i][next.j], _ = cur.at(steps)
			steps[next.i][next.j]++
			Q = append(Q, next)
		}

	}
}
func buildRoad(steps [][]int) []point {
	lastPoint := point{len(steps) - 1, len(steps[0]) - 1}
	lastStep := steps[lastPoint.i][lastPoint.j]
	road := make([]point, lastStep+1, lastStep+1)
	road[lastStep] = lastPoint
	for lastStep > 0 {
		lastStep--
		for _, dir := range directions {
			next := lastPoint.add(&dir)
			if val, ok := next.at(steps); ok && val == lastStep {
				road[lastStep] = next
				lastPoint = next
				break
			}
		}
	}
	return road
}
func main() {

	maze, steps := readMaze("learngo/maze/maze.in")
	walk(maze, steps)
	fmt.Println(buildRoad(steps))
}
