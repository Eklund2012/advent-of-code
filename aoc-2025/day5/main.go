package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(file string) ([][2]int, []int) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var (
		ranges [][2]int
		nums   []int
	)
	readingRanges := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" { // means start reading IDs
			readingRanges = false
			continue
		}
		if readingRanges {
			parts := strings.Split(line, "-")
			lo, _ := strconv.Atoi(parts[0])
			hi, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, [2]int{lo, hi})
		} else {
			n, _ := strconv.Atoi(line)
			nums = append(nums, n)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return ranges, nums
}

func checkInRange(nums []int, ranges [][2]int) int {
	fresh := 0
	for _, n := range nums {
		for _, r := range ranges {
			if n >= r[0] && n <= r[1] {
				fresh++
				break
			}
		}
	}
	return fresh
}

func part2(ranges [][2]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	var merged [][2]int
	current := ranges[0]
	for i := 1; i < len(ranges); i++ {
		next := ranges[i]

		if next[0] <= current[1]+1 {
			if next[1] > current[1] {
				current[1] = next[1]
			}
		} else {
			merged = append(merged, current)
			current = next
		}
	}
	merged = append(merged, current)
	total := 0
	for _, r := range merged {
		total += r[1] - r[0] + 1
	}
	return total
}

func main() {
	file := "input.txt"

	ranges, nums := part1(file)
	fresh := checkInRange(nums, ranges)
	fmt.Printf("%d Are fresh (part1)\n", fresh)
	numFreshIDs := part2(ranges)
	fmt.Printf("%d ingredient IDs are considered fresh (part2)", numFreshIDs)
}
