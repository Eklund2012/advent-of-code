package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		for _, n := range nums {
			fmt.Println(n)
		}
	}
}

func main() {
	file := "test.txt"
	part1(file)
}
