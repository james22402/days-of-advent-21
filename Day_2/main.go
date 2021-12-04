package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	part_one()
	println("")
	part_two()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one() {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	horizontal_pos := 0
	depth := 0

	for scanner.Scan() {
		instr := strings.Split(scanner.Text(), " ")
		switch instr[0] {
		case "forward":
			pos, err := strconv.Atoi(instr[1])
			check(err)
			horizontal_pos += pos
		case "up":
			pos, err := strconv.Atoi(instr[1])
			check(err)
			depth -= pos
		case "down":
			pos, err := strconv.Atoi(instr[1])
			check(err)
			depth += pos
		}
	}
	println("HZ Pos: " + strconv.Itoa(horizontal_pos))
	println("Depth: " + strconv.Itoa(depth))
	println("Answer: " + strconv.Itoa(horizontal_pos*depth))
}

func part_two() {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	horizontal_pos := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		instr := strings.Split(scanner.Text(), " ")
		switch instr[0] {
		case "forward":
			pos, err := strconv.Atoi(instr[1])
			check(err)
			horizontal_pos += pos
			depth += aim * pos
		case "up":
			pos, err := strconv.Atoi(instr[1])
			check(err)
			aim -= pos
		case "down":
			pos, err := strconv.Atoi(instr[1])
			check(err)
			aim += pos
		}
	}
	println("HZ Pos: " + strconv.Itoa(horizontal_pos))
	println("Depth: " + strconv.Itoa(depth))
	println("Answer: " + strconv.Itoa(horizontal_pos*depth))
}
