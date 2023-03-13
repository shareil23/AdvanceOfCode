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

type coord struct {
  x, y int
}

func solution() int {
  // code here
  var the_maze [][]string
  var temp_array []string
  var s_position, e_position coord
  var is_e_find bool = true
  var step int

  // got the maze
  for x_index, lines := range readFiles() {
    temp_array = nil
    for y_index, val := range lines {
      // locate position of S
      if string(val) == "S" {
        s_position.x = x_index
        s_position.y = y_index
      }

      // locate position of E
      if string(val) == "E" {
        e_position.x = x_index
        e_position.y = y_index
      }

      temp_array = append(temp_array, string(val))
    }
    the_maze = append(the_maze, temp_array)
  }

  // start looping to find E
  for is_e_find {
    step++

    // check is E position finded
    if step > 99 {
      is_e_find = false
    }
  }

  fmt.Println(the_maze)
  fmt.Println(s_position)
  fmt.Println(e_position)

  return 0
}


func main() {
  // for _, lines := range readFiles() {
      // the code here
  // }
  fmt.Println(solution())
}
