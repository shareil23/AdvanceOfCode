package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	init_container = [][]string{
		{"[H]", "[B]", "[V]", "[W]", "[N]", "[M]", "[L]", "[P]"}, // container no 1
		{"[M]", "[Q]", "[H]"}, // container no 2
		{"[N]", "[D]", "[B]", "[G]", "[F]", "[Q]", "[M]", "[L]"}, // container no 3
		{"[Z]", "[T]", "[F]", "[Q]", "[M]", "[W]", "[G]"},        // container no 4
		{"[M]", "[T]", "[H]", "[P]"},                             // container no 5
		{"[C]", "[B]", "[M]", "[J]", "[D]", "[H]", "[G]", "[T]"}, // container no 6
		{"[M]", "[N]", "[B]", "[F]", "[V]", "[R]"},               // container no 7
		{"[P]", "[L]", "[H]", "[M]", "[R]", "[G]", "[S]"},        // container no 8
		{"[P]", "[D]", "[B]", "[C]", "[N]"},                      // container no 9
	}
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

func solution() string {
	// code here
	var result string

	for _, lines := range readFiles() {
		/*
		   Details for the line
		   will be splited by ","
		   [0] is total data need to move
		   [1] from index
		   [2] target index
		   the index should be -1 due array start at 0
		*/
		lines_data := strings.Split(lines, ",")
		total_data, _ := strconv.Atoi(lines_data[0])
		selected_index, _ := strconv.Atoi(lines_data[1])
		targeted_index, _ := strconv.Atoi(lines_data[2])
		current_index := len(init_container[selected_index-1]) - total_data

		// create temporaryArray
		var temp_array_1, temp_array_2 []string

		// copy selected index data to temp_array_1
		temp_array_1 = append(
			temp_array_1,
			init_container[selected_index-1][current_index:]...,
		)

		// insert new ordered data
		for i := len(temp_array_1) - 1; i >= 0; i-- {
			temp_array_2 = append(temp_array_2, temp_array_1[i])
		}

		// move the total data from selcted index to target index
		init_container[targeted_index-1] = append(
			init_container[targeted_index-1][0:],
			temp_array_2...,
		)

		// remove the moved element from target index
		init_container[selected_index-1] = append(
			init_container[selected_index-1][:0],
			init_container[selected_index-1][:current_index]...,
		)

		fmt.Println(init_container)
	}

	for _, i := range init_container {
		result += i[len(i)-1]
	}

	result = strings.ReplaceAll(result, "[", "")
	result = strings.ReplaceAll(result, "]", "")

	return result
}

func solution2() string {
	// code here
	var result string

	for _, lines := range readFiles() {
		/*
		   Details for the line
		   will be splited by ","
		   [0] is total data need to move
		   [1] from index
		   [2] target index
		   the index should be -1 due array start at 0
		*/
		lines_data := strings.Split(lines, ",")
		total_data, _ := strconv.Atoi(lines_data[0])
		selected_index, _ := strconv.Atoi(lines_data[1])
		targeted_index, _ := strconv.Atoi(lines_data[2])
		current_index := len(init_container[selected_index-1]) - total_data

		// move the total data from selcted index to target index
		init_container[targeted_index-1] = append(
			init_container[targeted_index-1][0:],
			init_container[selected_index-1][current_index:]...,
		)

		// remove the moved element from target index
		init_container[selected_index-1] = append(
			init_container[selected_index-1][:0],
			init_container[selected_index-1][:current_index]...,
		)
	}

	for _, i := range init_container {
		result += i[len(i)-1]
	}

	result = strings.ReplaceAll(result, "[", "")
	result = strings.ReplaceAll(result, "]", "")

	return result
}

func main() {
	fmt.Println(solution2())
}
