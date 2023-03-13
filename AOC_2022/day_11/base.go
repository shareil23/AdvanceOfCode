package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

type monkey struct {
	code            int
	items           []int
	operation_math  string
	operation_value string
	test_value      int
	is_true         int
	is_false        int
}

func solution() int {
	// code here
	var monkeys []monkey
	var count_monkey int
	var inspected_item = make(map[int]int)
	var result []int

	for _, lines := range readFiles() {
		// the code here
		if strings.Contains(lines, "Monkey") {
			monkeys = append(monkeys, monkey{code: count_monkey})
			inspected_item[count_monkey] = 0
			continue
		}

		temp_string := strings.Split(lines, ":")

		if strings.Contains(temp_string[0], "Starting") {
			// remove space from string
			temp_string[1] = strings.ReplaceAll(temp_string[1], " ", "")
			temp_val := strings.Split(temp_string[1], ",")
			var temp_slice []int

			for _, i := range temp_val {
				temp_i_value, _ := strconv.Atoi(i)
				temp_slice = append(temp_slice, temp_i_value)
			}

			// insert new data to struct
			monkeys[count_monkey].items = append(monkeys[count_monkey].items, temp_slice...)
		} else if strings.Contains(temp_string[0], "Operation") {
			// remove the "new = " stirng
			temp_string[1] = strings.ReplaceAll(temp_string[1], " new = ", "")
			temp_val := strings.Split(temp_string[1], " ")

			monkeys[count_monkey].operation_math = temp_val[1]
			monkeys[count_monkey].operation_value = temp_val[len(temp_val)-1]
		} else if strings.Contains(temp_string[0], "Test") {
			temp_val := strings.Split(temp_string[1], " ")

			monkeys[count_monkey].test_value, _ = strconv.Atoi(temp_val[len(temp_val)-1])
		} else if strings.Contains(temp_string[0], "true") {
			temp_val := strings.Split(temp_string[1], " ")

			monkeys[count_monkey].is_true, _ = strconv.Atoi(temp_val[len(temp_val)-1])
		} else if strings.Contains(temp_string[0], "false") {
			temp_val := strings.Split(temp_string[1], " ")

			monkeys[count_monkey].is_false, _ = strconv.Atoi(temp_val[len(temp_val)-1])
		}

		// if the lines equal empty string then the count_monkey increment
		if lines == "" {
			count_monkey++
		}
	}

	// loop the round
	round := 0

	for round != 20 {
		// loop the monkey
		for index, i := range monkeys {
			// loop the stress level
			if i.items != nil {
				for _, j := range i.items {
					temp_val, _ := strconv.Atoi(i.operation_value)
					temp_sum := 0

					if i.operation_value == "old" {
						temp_val = j
					}

					switch i.operation_math {
					case "+":
						temp_sum = j + temp_val
					case "-":
						temp_sum = j - temp_val
					case "*":
						temp_sum = j * temp_val
					case "/":
						temp_sum = j / temp_val
					}

					// do a test phase
					temp_sum_divide := math.Floor(float64(temp_sum) / 3)

					// println("Code: ", i.code, " item: ", j, " divided result: ", int(temp_sum_divide), " the test result: ", int(temp_sum_divide)%i.test_value == 0)

					// check where need to throw the item
					if int(temp_sum_divide)%i.test_value == 0 {
						monkeys[i.is_true].items = append(monkeys[i.is_true].items, int(temp_sum_divide))
					} else {
						monkeys[i.is_false].items = append(monkeys[i.is_false].items, int(temp_sum_divide))
					}

					inspected_item[i.code] += 1
				}

				monkeys[index].items = nil
			}
		}

		// fmt.Println(monkeys)
		round++
	}

	result = sortMap(inspected_item)
	return result[0] * result[1]
}

func sortMap(inspected_item map[int]int) []int {
	keys := make([]int, len(inspected_item))

	for _, k := range inspected_item {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	return keys
}

func solution2() int {
	// code here
	var monkeys []monkey
	var count_monkey int
	var inspected_item = make(map[int]int)
	var result []int

	for _, lines := range readFiles() {
		// the code here
		if strings.Contains(lines, "Monkey") {
			monkeys = append(monkeys, monkey{code: count_monkey})
			inspected_item[count_monkey] = 0
			continue
		}

		temp_string := strings.Split(lines, ":")

		if strings.Contains(temp_string[0], "Starting") {
			// remove space from string
			temp_string[1] = strings.ReplaceAll(temp_string[1], " ", "")
			temp_val := strings.Split(temp_string[1], ",")
			var temp_slice []int

			for _, i := range temp_val {
				temp_i_value, _ := strconv.Atoi(i)
				temp_slice = append(temp_slice, temp_i_value)
			}

			// insert new data to struct
			monkeys[count_monkey].items = append(monkeys[count_monkey].items, temp_slice...)
		} else if strings.Contains(temp_string[0], "Operation") {
			// remove the "new = " stirng
			temp_string[1] = strings.ReplaceAll(temp_string[1], " new = ", "")
			temp_val := strings.Split(temp_string[1], " ")

			monkeys[count_monkey].operation_math = temp_val[1]
			monkeys[count_monkey].operation_value = temp_val[len(temp_val)-1]
		} else if strings.Contains(temp_string[0], "Test") {
			temp_val := strings.Split(temp_string[1], " ")

			monkeys[count_monkey].test_value, _ = strconv.Atoi(temp_val[len(temp_val)-1])
		} else if strings.Contains(temp_string[0], "true") {
			temp_val := strings.Split(temp_string[1], " ")

			monkeys[count_monkey].is_true, _ = strconv.Atoi(temp_val[len(temp_val)-1])
		} else if strings.Contains(temp_string[0], "false") {
			temp_val := strings.Split(temp_string[1], " ")

			monkeys[count_monkey].is_false, _ = strconv.Atoi(temp_val[len(temp_val)-1])
		}

		// if the lines equal empty string then the count_monkey increment
		if lines == "" {
			count_monkey++
		}
	}

	// find number to manage stress level
  // with mod method (a+b) mod x
	mod := 1

	for _, i := range monkeys {
		mod *= i.test_value
	}

	// loop the round
	round := 0

	for round != 10000 {
		// loop the monkey
		for index, i := range monkeys {
			// loop the stress level
			if i.items != nil {
				for _, j := range i.items {
					temp_val, _ := strconv.Atoi(i.operation_value)
					temp_sum := 0

					if i.operation_value == "old" {
						temp_val = j
					}

					switch i.operation_math {
					case "+":
						temp_sum = j + temp_val
					case "-":
						temp_sum = j - temp_val
					case "*":
						temp_sum = j * temp_val
					case "/":
						temp_sum = j / temp_val
					}

					// do a test phase
					temp_sum %= mod

					// println("Code: ", i.code, " item: ", j, " divided result: ", int(temp_sum_divide), " the test result: ", int(temp_sum_divide)%i.test_value == 0)

					// check where need to throw the item
					if temp_sum%i.test_value == 0 {
						monkeys[i.is_true].items = append(monkeys[i.is_true].items, int(temp_sum))
					} else {
						monkeys[i.is_false].items = append(monkeys[i.is_false].items, int(temp_sum))
					}

					inspected_item[i.code] += 1
				}

				monkeys[index].items = nil
			}
		}

		// fmt.Println(monkeys)
		round++
	}

	result = sortMap(inspected_item)
	fmt.Println(inspected_item)
	return result[0] * result[1]
}

func main() {
	fmt.Println(solution2())
}
