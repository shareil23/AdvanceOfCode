package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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


func solution(inputs []string) int {
  var result int
  var temp_sum int

  for _, lines := range inputs {
    if lines == "" {
      if temp_sum > result {
        result = temp_sum
      }

      temp_sum = 0
    } else {
      parse_int, _ := strconv.Atoi(lines)
      temp_sum += parse_int
    }
  }

  if temp_sum > result {
    result = temp_sum
  }

  return result
}


func solution2(inputs []string) int {
  var result int
  var list_of_result []int
  var temp_sum int

  for _, lines := range inputs {
    if lines == "" {
      if temp_sum > result {
        result = temp_sum
      }

      list_of_result = append(list_of_result, temp_sum)
      temp_sum = 0
    } else {
      parse_int, _ := strconv.Atoi(lines)
      temp_sum += parse_int
    }
  }

  list_of_result = append(list_of_result, temp_sum)

  if temp_sum > result {
    result = temp_sum
  }

  // sort the list_of_result
  sort.Sort(sort.Reverse(sort.IntSlice(list_of_result)))

  fmt.Println(list_of_result)
  result = 0
  for index, value := range list_of_result{
    if index < 3 {
      result += value
    }

    if index == 2 {
      break
    }
  }

  return result
}


func main() {
  fmt.Println(solution2(readFiles()))
}
