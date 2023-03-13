package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFiles() []string {
	file, err := os.Open("test.txt")

	// validation if error
	if err != nil {
		log.Fatalf("Failed to open.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	file.Close()

	return result
}

func solution() int {
	// code here
	var cycles int
	var x int = 1
	var cycles_target int = 20
	var result int

	for _, lines := range readFiles() {
		/**
		result_data[0] determine the signal addx or noop
		result_data[1] determine the signal value
		*/
		result_data := strings.Split(lines, " ")
		temp_counter := 1
		temp_value_data := 0

		if result_data[0] == "addx" {
			temp_counter = 2
			value_data, _ := strconv.Atoi(result_data[1])
			temp_value_data = value_data
		}

		// loop the cycle
		for temp_counter > 0 {
			cycles += 1

			if cycles >= cycles_target {
				result += cycles * x
				cycles_target += 40
			}
			temp_counter--
		}

		x += temp_value_data
	}

	return result
}

func solution2() string {
	// code here
	var cycles int
	var x int = 1
	var result string
	var curent_pos_sprite int = 0
	var curent_pos_crt int = 0

	for _, lines := range readFiles() {
		/**
		result_data[0] determine the signal addx or noop
		result_data[1] determine the signal value
		*/
		result_data := strings.Split(lines, " ")
		temp_counter := 1
		temp_value_data := 0

		if result_data[0] == "addx" {
			temp_counter = 2
			value_data, _ := strconv.Atoi(result_data[1])
			temp_value_data = value_data
		}

		// loop the cycle
		for temp_counter > 0 {
			cycles += 1

			if curent_pos_crt >= 40 {
				curent_pos_crt = 0
				result += "\n"
			}

      // to check the curernt pos of crt are in between range pos sprite
			if curent_pos_crt >= curent_pos_sprite-1 && curent_pos_crt <= curent_pos_sprite+1 {
				result += "#"
			} else {
				result += "."
			}

			curent_pos_crt += 1
			temp_counter--
		}

		x += temp_value_data
		curent_pos_sprite = x

  }

	return result
}

func main() {
	fmt.Println(solution2())
}
