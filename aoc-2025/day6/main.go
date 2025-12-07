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
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // do NOT trim spaces
	}
	return lines, scanner.Err()
}

func padLines(lines []string) []string {
	maxLen := 0
	for _, l := range lines {
		if len(l) > maxLen {
			maxLen = len(l)
		}
	}

	out := make([]string, len(lines))
	for i, l := range lines {
		if len(l) < maxLen {
			l = l + strings.Repeat(" ", maxLen-len(l))
		}
		out[i] = l
	}
	return out
}

func isSpaceColumn(lines []string, col int) bool {
	for _, row := range lines {
		if row[col] != ' ' {
			return false
		}
	}
	return true
}

func detectBlocks(lines []string) [][2]int {
	width := len(lines[0])
	var blocks [][2]int
	inBlock := false
	var start int

	for col := 0; col < width; col++ {
		if isSpaceColumn(lines, col) {
			if inBlock {
				blocks = append(blocks, [2]int{start, col - 1})
				inBlock = false
			}
		} else {
			if !inBlock {
				start = col
				inBlock = true
			}
		}
	}
	if inBlock {
		blocks = append(blocks, [2]int{start, width - 1})
	}
	return blocks
}

func extractNumbers(row string, blocks [][2]int) ([]int, []string) {
	var exNums []int
	var exOps []string

	for _, block := range blocks {
		start, end := block[0], block[1]
		sub := strings.TrimSpace(row[start : end+1])

		if sub == "" {
			continue
		}

		if sub == "+" || sub == "*" {
			exOps = append(exOps, sub)
		} else {
			v, _ := strconv.Atoi(sub)
			exNums = append(exNums, v)
		}
	}
	return exNums, exOps
}

func main() {
	file := "test.txt"
	grid, _ := readGrid(file)
	grid = padLines(grid)
	blocks := detectBlocks(grid)

	problems := make([][]int, len(blocks))
	problemOps := make([][]string, len(blocks))
	for i, row := range grid {
		problems[i], problemOps[i] = extractNumbers(row, blocks)
	}
	fmt.Println(problems, problemOps)

}
