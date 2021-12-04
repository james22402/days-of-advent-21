package main

import (
	"bufio"
	"os"
	"strconv"
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
	// Hello world, the web server
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	current := 0
	i := 0
	prev := 0

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		current, err = strconv.Atoi(scanner.Text())
		check(err)
		if current > prev {
			if prev != 0 {
				i++
			}
		}
		prev, err = strconv.Atoi(scanner.Text())
		check(err)
	}
	print(i)
	file.Close()

}

func part_two() {
	// Hello world, the web server
	file, err := os.Open("input.txt")
	check(err)

	window := make([]int, 0)
	window_depth := 0
	prev_window := make([]int, 0)
	prev_depth := 0
	i := 0
	j := 0

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		num1, err := strconv.Atoi(scanner.Text())
		window = append(window, num1)
		check(err)
		if j >= 1 {
			window_depth = sum(window)
			prev_depth = sum(prev_window)
			if window_depth > prev_depth {
				i++
			}
		} else {
			scanner.Scan()
			num2, err := strconv.Atoi(scanner.Text())
			check(err)
			window = append(window, num2)
			scanner.Scan()
			num3, err := strconv.Atoi(scanner.Text())
			window = append(window, num3)
			check(err)
		}
		prev_window = window
		window = window[1:]
		j++
	}
	print(i)
	file.Close()
}

func sum(array []int) int {
	sum := 0
	for _, s := range array {
		sum += s
	}
	return sum
}
