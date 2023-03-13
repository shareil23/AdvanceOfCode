package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func getAlphabetIndex(character string) int {
	var char_to_rune rune = []rune(strings.ToLower(character))[0]

	return int(char_to_rune - 'a' + 1)
}

func getAlphabetIndex2(character string) int {
	var alphabets string = ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	return strings.Index(alphabets, character)
}

func checkUpper(character string) bool {
	if strings.ToUpper(character) == character {
		return true
	}

	return false
}

func countSameCharacter(first_item, second_item string) string {
	var result = make(map[string]int)
	var temp_val []string
	var result_character string
	var temp_value_counter int = 0

	// init the first value
	for _, val := range first_item {
		result[string(val)] = 0
		temp_val = append(temp_val, string(val))
	}

	// check with the second value
	for _, val := range second_item {
		if _, is_ok := result[string(val)]; is_ok {
			result[string(val)] += 1
		} else {
			result[string(val)] = 0
		}
	}

	fmt.Println(result)
	// filter the biggest value
	for _, val := range temp_val {
		if result[val] > temp_value_counter {
			result_character = val
			temp_value_counter = result[val]
		}
	}

	fmt.Println(result_character)

	return result_character
}

func solution() int {
	var result int

	for _, lines := range readFiles() {
		var length_of_text int = len(lines) / 2
		var first_item string = lines[0:length_of_text]
		var second_item string = lines[length_of_text:]
		var result_item string = countSameCharacter(first_item, second_item)

		if checkUpper(result_item) {
			result += getAlphabetIndex(result_item) + 26
		} else {
			result += getAlphabetIndex(result_item)
		}
	}

	return result
}

func solution2() int {
	var result [][]int
	var temp_result []int
	var counter int
	var temp_char_container []string
	var result_sum int

	for _, lines := range readFiles() {
		for _, values := range lines {
			// check the array contain the character
			if !contains(temp_char_container, string(values)) {
				temp_result = append(temp_result, getAlphabetIndex2(string(values)))
				temp_char_container = append(temp_char_container, string(values))
			}
		}

		result = append(result, temp_result)
		temp_result = nil
		temp_char_container = nil
		counter++

		if counter < 3 {
			continue
		}

		phase_1_data := retainAllSlice(result[0], result[1])
		phase_2_data := retainAllSlice(phase_1_data, result[2])

		// sum the final data phase_2_data
		for _, i := range phase_2_data {
			result_sum += i
		}

		// reset value
		result = nil
		counter = 0
	}

	return result_sum
}

func retainAllSlice(a, b []int) []int {
	// a := the data, b := the comparator, the result remove the element on slice a that no contain value in slice b
	var result []int
	for _, i := range a {
		if containsInt(b, i) {
			result = append(result, i)
		}
	}

	return result
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(solution2())
}
