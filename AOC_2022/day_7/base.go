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
	var dir_temp []string
	var total_size int
	var result = make(map[string]int)

	// init root dir
	result["/"] = 0

	for _, lines := range readFiles() {
		temp_string := strings.Split(lines, " ")

		if len(temp_string) >= 3 {
			if temp_string[1] == "cd" && temp_string[2] != ".." {
				dir_temp = append(dir_temp, temp_string[2])
			} else {
				// to append value sub dir to root dir
				if len(dir_temp) > 2 {
					result[concatFullPath(dir_temp[:len(dir_temp)-1])] += result[concatFullPath(dir_temp[:len(dir_temp)])]
				}

				// remove the last element from slice
				dir_temp = dir_temp[:len(dir_temp)-1]
			}
		} else {
			// if the value not equal dir but file size information then will append to map
			if temp_string[0] != "dir" {
				temp_parse, _ := strconv.Atoi(temp_string[0])
				result[concatFullPath(dir_temp)] += temp_parse
			}
		}
	}

	// recalculate the size of root
	for key, value := range result {
		if strings.Count(key, "/") == 1 && key != "/" {
			result["/"] += value
		}

		if value <= 100000 {
			total_size += value
		}
	}

	return total_size
}

func concatFullPath(dir_temp []string) string {
	var result string

	for index, value := range dir_temp {
		if index < 2 {
			result += value
		} else {
			result += "/" + value
		}
	}

	return result
}

func solution2() int {
	// code here
	var dir_temp []string
	var total_size int
	var result = make(map[string]int)

	// init root dir
	result["/"] = 0

	for _, lines := range readFiles() {
		temp_string := strings.Split(lines, " ")

		if len(temp_string) >= 3 {
			if temp_string[1] == "cd" && temp_string[2] != ".." {
				dir_temp = append(dir_temp, temp_string[2])
			} else {
				// to append value sub dir to root dir
				if len(dir_temp) > 2 {
					result[concatFullPath(dir_temp[:len(dir_temp)-1])] += result[concatFullPath(dir_temp[:len(dir_temp)])]
				}

				// remove the last element from slice
				dir_temp = dir_temp[:len(dir_temp)-1]
			}
		} else {
			// if the value not equal dir but file size information then will append to map
			if temp_string[0] != "dir" {
				temp_parse, _ := strconv.Atoi(temp_string[0])
				result[concatFullPath(dir_temp)] += temp_parse
			}
		}
	}

	// recalculate the size of root
	for key, value := range result {
		if strings.Count(key, "/") == 1 && key != "/" {
			result["/"] += value
		}

		if value <= 100000 {
			total_size += value
		}
	}

	var total_space int = 70000000
	var space_required int = 30000000
	var space_needed int = result["/"] - (total_space - space_required)
	total_size = total_space

	for _, value := range result {
		if value >= space_needed {
			if value < total_size {
				total_size = value
			}
		}
	}

	return total_size
}

func main() {
	fmt.Println(solution2())
}
