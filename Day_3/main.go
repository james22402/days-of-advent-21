package main

import (
	"bufio"
	"fmt"
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
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	pos_ray_ones := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for scanner.Scan() {
		for i, c := range scanner.Text() {
			if c == '1' {
				pos_ray_ones[i]++
			}
		}
	}
	gamma_output := ""
	epsilon_output := ""
	for _, count := range pos_ray_ones {
		if count >= 500 {
			gamma_output += "1"
			epsilon_output += "0"
		} else {
			gamma_output += "0"
			epsilon_output += "1"
		}
	}
	gamma_num, err := strconv.ParseInt(gamma_output, 2, 64)
	check(err)
	epsilon_num, err := strconv.ParseInt(epsilon_output, 2, 64)
	check(err)
	println("G: " + gamma_output + "   " + fmt.Sprint(gamma_num))
	println("E: " + epsilon_output + "   " + fmt.Sprint(epsilon_num))
	println("Answer: " + fmt.Sprint(gamma_num*epsilon_num))

}

func part_two() {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
}
