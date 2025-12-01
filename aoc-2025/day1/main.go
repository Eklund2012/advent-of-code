package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currVal := 50
	password := 0

	for scanner.Scan() {
		line := scanner.Text()

		dir := string(line[0])

		n, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(err)
		}

		if dir == "L" {
			currVal = (currVal - n + 100) % 100
		} else {
			currVal = (currVal + n) % 100
		}

		if currVal == 0 {
			password += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(password)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currVal := 50
	password := 0

	for scanner.Scan() {
		line := scanner.Text()

		dir := string(line[0])

		n, err := strconv.Atoi(line[1:])

		fullRotations := int(math.Abs(float64(n / 100)))
		n = n % 100

		password += fullRotations

		if err != nil {
			panic(err)
		}

		if dir == "L" {
			if n > currVal {
				if currVal != 0 {
					password++
				}
				currVal = (currVal - n) + 100
			} else {
				currVal = (currVal - n)
			}
		} else {
			if (currVal + n) > 100 {
				password++
			}
			currVal = (currVal + n) % 100
		}

		if currVal == 0 {
			password += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(password)
}

func main() {
	part1()
	part2()
}
