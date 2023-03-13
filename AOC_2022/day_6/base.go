package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func IsStringUnique(concat_string string) bool {
  var temp_map = make(map[string]int)

  for _, i := range concat_string {
    if _, is_ok := temp_map[string(i)]; is_ok {
      return false
    } else {
      temp_map[string(i)] += 1
    }
  }

  return true
}

func solution() int {
  // code here
  var target_index int = 4
  var result int

  for _, lines := range readFiles() {
    for index, _ := range string(lines) {
      if target_index > len(lines) {
        break
      }

      if IsStringUnique(lines[index:target_index]){
        result = target_index
        fmt.Println("result: ", result)
        break
      }
      target_index += 1
    }

    target_index = 4
  }

  return result
}

func solution2() int {
  // code here
  var target_index int = 14
  var result int

  for _, lines := range readFiles() {
    for index, _ := range string(lines) {
      if target_index > len(lines) {
        break
      }

      if IsStringUnique(lines[index:target_index]){
        result = target_index
        fmt.Println("result: ", result)
        break
      }
      target_index += 1
    }

    target_index = 14
  }

  return result
}

func main() {
  // for _, lines := range readFiles() {
      // the code here
  // }
  fmt.Println(solution2())
}
