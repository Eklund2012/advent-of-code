package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readGrid(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() { // Read the grid
		lines = append(lines, scanner.Text()) // keep spaces
	}
	return lines, scanner.Err()
}

func pad(lines []string) []string {
	maxLen := 0
	for _, l := range lines { // find longest line
		if len(l) > maxLen {
			maxLen = len(l)
		}
	}
	out := make([]string, len(lines))
	for i, l := range lines {
		diff := maxLen - len(l)
		if diff > 0 {
			l += strings.Repeat(" ", diff) // add spaces
		}
		out[i] = l
	}
	return out
}

func isSpaceColumn(lines []string, col int) bool { // check if a column is only spaces -> seperator column
	for _, row := range lines {
		if row[col] != ' ' {
			return false
		}
	}
	return true
}

func findBlocks(lines []string) [][2]int {
	w := len(lines[0])
	var blocks [][2]int

	in := false
	start := 0

	for c := range w {
		if isSpaceColumn(lines, c) {
			if in {
				blocks = append(blocks, [2]int{start, c - 1})
				in = false
			}
		} else if !in {
			in = true
			start = c
		}
	}
	if in {
		blocks = append(blocks, [2]int{start, w - 1})
	}

	return blocks
}

func main() {
	grid, _ := readGrid("input.txt")
	grid = pad(grid)
	blocks := findBlocks(grid)

	numProblems := len(blocks)
	problems := make([][]int, numProblems)
	ops := make([]string, numProblems)

	last := len(grid) - 1

	for r, row := range grid {
		for i, b := range blocks {
			sub := strings.TrimSpace(row[b[0] : b[1]+1])

			if r == last {
				ops[i] = sub
				continue
			}

			if sub == "" {
				continue
			}

			v, _ := strconv.Atoi(sub)
			problems[i] = append(problems[i], v)
		}
	}

	total := 0
	for i := range problems {
		vals := problems[i]
		op := ops[i]

		acc := vals[0]
		for _, v := range vals[1:] {
			if op == "+" {
				acc += v
			} else {
				acc *= v
			}
		}
		total += acc
	}

	fmt.Println("TOTAL =", total)
}
