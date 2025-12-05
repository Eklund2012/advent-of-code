package main

import (
	"bufio"
	"fmt"
	"os"
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
		if line == "" { // means we start reading IDs
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

func main() {
	file := "input.txt"

	ranges, nums := part1(file)
	fresh := checkInRange(nums, ranges)
	fmt.Printf("%d Are fresh (part1)\n", fresh)

}
