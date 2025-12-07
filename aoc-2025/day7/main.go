package main

import (
	"bufio"
	"fmt"
	"os"
)

func findStart(grid [][]rune) (int, int) {
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 'S' {
				return r, c
			}
		}
	}
	return -1, -1
}

type Beam struct {
	r, c int
}

func splitAnalyzer(grid [][]rune, beamGrid [][]rune, queue []Beam) int {
	splits := 0

	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:]
		r, c := b.r, b.c
		r++
		if r >= len(grid) {
			continue
		}
		if grid[r][c] == '.' {
			beamGrid[r][c] = '|'
			queue = append(queue, Beam{r, c})
		} else if grid[r][c] == '^' {
			if beamGrid[r][c] != '^' {
				beamGrid[r][c] = '^'
				splits++
			}
			if c-1 >= 0 && beamGrid[r][c-1] != '|' {
				beamGrid[r][c-1] = '|'
				queue = append(queue, Beam{r, c - 1})

			}
			if c+1 < len(grid[0]) && beamGrid[r][c+1] != '|' {
				beamGrid[r][c+1] = '|'
				queue = append(queue, Beam{r, c + 1})
			}
		}
	}
	return splits
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	sr, sc := findStart(grid)
	fmt.Println("Start: ", sr, sc)

	beamGrid := make([][]rune, len(grid))
	for i := range beamGrid {
		beamGrid[i] = make([]rune, len(grid[0]))
		for j := range beamGrid[i] {
			beamGrid[i][j] = '.' // empty space
		}
	}

	beamGrid[sr][sc] = 'S'
	var queue []Beam
	queue = append(queue, Beam{sr, sc})

	splits := splitAnalyzer(grid, beamGrid, queue)

	fmt.Println("Final grid:")
	for _, row := range beamGrid {
		fmt.Println(string(row))
	}
	fmt.Println("Number of splits: ", splits)
}
