package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	value := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstIndex := 0

		for i := 1; i < len(line)-1; i++ {
			val := int(line[i] - '0')
			if val > int(line[firstIndex]-'0') {
				firstIndex = i
			}
		}
		secIndex := firstIndex + 1
		for i := secIndex; i < len(line); i++ {
			val := int(line[i] - '0')
			if val > int(line[secIndex]-'0') {
				secIndex = i
			}
		}
		digit1 := int(line[firstIndex] - '0')
		digit2 := int(line[secIndex] - '0')
		value += digit1*10 + digit2
	}
	fmt.Println("Total output joltage part1: ", value)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	value := 0
	for scanner.Scan() {
		line := scanner.Text()
		stack := []int{}
		remaining := len(line)

		for _, d := range line {
			digit := int(d - '0')
			for len(stack) > 0 && stack[len(stack)-1] < digit && len(stack)-1+remaining >= 12 {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, digit)
			remaining--
		}
		if len(stack) > 12 {
			stack = stack[:12]
		}
		num := 0
		for _, d := range stack {
			num = num*10 + d
		}
		fmt.Println(num)
		value += num
	}
	fmt.Println("Total output joltage part2: ", value)
}

func main() {
	part1()
	part2()
}
