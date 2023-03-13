package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	var trees [][]int
	var total_tree_outside int

	// move from text to 2d array
	// the trees 0 is smallest and 9 is tallest
	for _, lines := range readFiles() {
		var temp_array []int

		for _, i := range lines {
			// add the data to matrix slice
			temp_parse, _ := strconv.Atoi(string(i))
			temp_array = append(temp_array, temp_parse)
		}

		trees = append(trees, temp_array)
	}

	// count visible trees
	total_tree_outside = (len(trees[0]) * 2) + (len(trees)-2)*2

	for index_i, i := range trees {
		if index_i == 0 || index_i == len(trees)-1 {
			continue
		}

		for index_j, _ := range i {
			if index_j == 0 || index_j == len(i)-1 {
				continue
			}

			if visibleChecker(trees, index_i, index_j) {
				total_tree_outside += 1
			}
		}
	}

	return total_tree_outside
}

func visibleChecker(data_tree [][]int, i_index, j_index int) bool {
	var (
		top_status    bool = true
		left_status   bool = true
		right_status  bool = true
		bottom_status bool = true
	)

	// check the current index tree is visible from top
	for i := i_index - 1; i >= 0; i-- {
		if data_tree[i][j_index] >= data_tree[i_index][j_index] {
			top_status = false
			break
		}
	}

	// check the current index tree is visible from left
	for j := j_index - 1; j >= 0; j-- {
		if data_tree[i_index][j] >= data_tree[i_index][j_index] {
			left_status = false
			break
		}
	}

	// check the current index tree is visible from right
	for j := j_index + 1; j <= len(data_tree[i_index])-1; j++ {
		if data_tree[i_index][j] >= data_tree[i_index][j_index] {
			right_status = false
			break
		}
	}

	// check the current index tree is visible from bottom
	for i := i_index + 1; i <= len(data_tree)-1; i++ {
		if data_tree[i][j_index] >= data_tree[i_index][j_index] {
			bottom_status = false
			break
		}
	}

	return top_status || left_status || right_status || bottom_status
}

func solution2() int {
	// code here
	var trees [][]int
	var total_tree_outside int

	// move from text to 2d array
	// the trees 0 is smallest and 9 is tallest
	for _, lines := range readFiles() {
		var temp_array []int

		for _, i := range lines {
			// add the data to matrix slice
			temp_parse, _ := strconv.Atoi(string(i))
			temp_array = append(temp_array, temp_parse)
		}

		trees = append(trees, temp_array)
	}

	for index_i, i := range trees {
		if index_i == 0 || index_i == len(trees)-1 {
			continue
		}

		for index_j, _ := range i {
			if index_j == 0 || index_j == len(i)-1 {
				continue
			}

			temp_cal := scenicChecker(trees, index_i, index_j)
			if temp_cal > total_tree_outside {
				total_tree_outside = temp_cal
			}
		}
	}

	return total_tree_outside
}

func scenicChecker(data_tree [][]int, i_index, j_index int) int {
	var (
		top_status    int
		left_status   int
		right_status  int
		bottom_status int
	)

	// check the current index tree is visible from top
	for i := i_index - 1; i >= 0; i-- {
		top_status += 1
		if data_tree[i][j_index] >= data_tree[i_index][j_index] {
			break
		}
	}

	// check the current index tree is visible from left
	for j := j_index - 1; j >= 0; j-- {
		left_status += 1
		if data_tree[i_index][j] >= data_tree[i_index][j_index] {
			break
		}
	}

	// check the current index tree is visible from right
	for j := j_index + 1; j <= len(data_tree[i_index])-1; j++ {
		right_status += 1
		if data_tree[i_index][j] >= data_tree[i_index][j_index] {
			break
		}
	}

	// check the current index tree is visible from bottom
	for i := i_index + 1; i <= len(data_tree)-1; i++ {
		bottom_status += 1
		if data_tree[i][j_index] >= data_tree[i_index][j_index] {
			break
		}
	}

	return top_status * left_status * right_status * bottom_status
}

func main() {
	fmt.Println(solution2())
}
