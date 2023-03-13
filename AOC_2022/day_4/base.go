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
	var result int

	for _, lines := range readFiles() {
		var temp_splited_strings []string = strings.Split(lines, ",")
		if checkSection(sectionToString(temp_splited_strings[0]), sectionToString(temp_splited_strings[1])) {
			result++
		}
	}

	return result
}

func sectionToString(section string) []int {
	var splited_string []string = strings.Split(section, "-")
	var result []int

	first_data, _ := strconv.Atoi(splited_string[0])
	second_data, _ := strconv.Atoi(splited_string[1])

	for i := first_data; i <= second_data; i++ {
    result = append(result, i)
	}

	return result
}

func checkSection(a, b []int) bool {
	var counter int
  var res_1, res_2 bool

	// check the first section a -> b
	for _, val := range a {
		for _, val2 := range b {
			if val2 == val {
				counter++
			}
		}
	}

	res_1 = len(a) == counter
	counter = 0

	// check the second section b -> a
	for _, val := range b {
		for _, val2 := range a {
			if val2 == val {
				counter++
			}
		}
	}

	res_2 = len(b) == counter
	return res_1 || res_2
}

func solution2() int {
	var result int

	for _, lines := range readFiles() {
		var temp_splited_strings []string = strings.Split(lines, ",")
		if checkSection2(sectionToString(temp_splited_strings[0]), sectionToString(temp_splited_strings[1])) {
			result++
		}
	}

	return result
}

func checkSection2(a, b []int) bool {
	var counter int
  var res_1, res_2 bool

	// check the first section a -> b
	for _, val := range a {
		for _, val2 := range b {
			if val2 == val {
				counter++
			}
		}
	}

	res_1 = counter != 0
	counter = 0

	// check the second section b -> a
	for _, val := range b {
		for _, val2 := range a {
			if val2 == val {
				counter++
			}
		}
	}

	res_2 = counter != 0
	return res_1 || res_2
}

func main() {
	fmt.Println(solution2())
}
