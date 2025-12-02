package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkInvalidIdP1(num int) bool {
	numStr := strconv.Itoa(num)

	mid := len(numStr) / 2
	if len(numStr)%2 == 0 {
		if numStr[0:mid] == numStr[mid:] {
			return true
		}
	}
	return false
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	line, _ := r.ReadString('\n')

	parts := strings.Split(strings.TrimSpace(line), ",")

	invalidIds := 0

	for _, p := range parts {
		bounds := strings.Split(p, "-")
		if len(bounds) != 2 {
			continue
		}

		num1, _ := strconv.Atoi(bounds[0])
		num2, _ := strconv.Atoi(bounds[1])

		counter := num1
		for counter <= num2 {
			if checkInvalidIdP1(counter) {
				invalidIds += counter
			}
			counter++
		}
	}
	fmt.Println("(part1) Invalid IDs: ", invalidIds)
}

func checkInvalidIdP2(num int) bool {
	numStr := strconv.Itoa(num)

	mid := len(numStr) / 2
	if len(numStr)%2 == 0 {
		if numStr[0:mid] == numStr[mid:] {
			return true
		}
	}

	for k := 1; k <= len(numStr)/2; k++ {
		if len(numStr)%k != 0 {
			continue
		}
		// check if repeating numStr[0:k] makes the whole string
		chunk := numStr[0:k]
		validRepeat := true
		for i := k; i < len(numStr); i += k {
			if chunk != numStr[i:i+k] {
				validRepeat = false
			}
		}
		if validRepeat {
			return true
		}
	}
	return false
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	line, _ := r.ReadString('\n')

	parts := strings.Split(strings.TrimSpace(line), ",")

	invalidIds := 0

	for _, p := range parts {
		bounds := strings.Split(p, "-")
		if len(bounds) != 2 {
			continue
		}

		num1, _ := strconv.Atoi(bounds[0])
		num2, _ := strconv.Atoi(bounds[1])

		counter := num1
		for counter <= num2 {
			if checkInvalidIdP2(counter) {
				invalidIds += counter
			}
			counter++
		}
	}
	fmt.Println("(part2) Invalid IDs: ", invalidIds)
}

func main() {
	part1()
	part2()
}
