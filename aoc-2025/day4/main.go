package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1() {
	dirs := [][2]int{
		{-1, 0},  // N
		{-1, 1},  // NE
		{0, 1},   // E
		{1, 1},   // SE
		{1, 0},   // S
		{1, -1},  // SW
		{0, -1},  // W
		{-1, -1}, // NW
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]string

	k := 0
	for scanner.Scan() {
		var row []string
		line := scanner.Text()
		for _, x := range line {
			row = append(row, string(x))
		}
		grid = append(grid, row)
		k++
	}
	rows := len(grid)
	cols := len(grid[0])
	rollsAccessed := 0
	for r := range rows {
		for c := range cols {
			adjacent := 0
			curr := grid[r][c]
			for _, d := range dirs {
				if curr != "@" {
					break
				}
				nr := r + d[0]
				nc := c + d[1]

				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}

				neighbor := grid[nr][nc]
				if curr == neighbor {
					adjacent++
				}
			}
			if adjacent < 4 && curr == "@" {
				rollsAccessed++
			}
		}
	}
	fmt.Println(rollsAccessed, " Rolls, of paper (part1)")
}

func part2() {
	dirs := [][2]int{
		{-1, 0},  // N
		{-1, 1},  // NE
		{0, 1},   // E
		{1, 1},   // SE
		{1, 0},   // S
		{1, -1},  // SW
		{0, -1},  // W
		{-1, -1}, // NW
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]string

	k := 0
	for scanner.Scan() {
		var row []string
		line := scanner.Text()
		for _, x := range line {
			row = append(row, string(x))
		}
		grid = append(grid, row)
		k++
	}
	rows := len(grid)
	cols := len(grid[0])
	rollsAccessed := 0
	changed := true
	for changed {
		changed = false
		for r := range rows {
			for c := range cols {
				adjacent := 0
				curr := grid[r][c]
				for _, d := range dirs {
					if curr != "@" {
						break
					}
					nr := r + d[0]
					nc := c + d[1]

					if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
						continue
					}

					neighbor := grid[nr][nc]
					if curr == neighbor {
						adjacent++
					}
				}
				if adjacent < 4 && curr == "@" {
					rollsAccessed++
					grid[r][c] = "."
					changed = true
				}
			}
		}
	}
	fmt.Println(rollsAccessed, " Rolls, of paper (part2)")
}

func main() {
	part1()
	part2()
}
