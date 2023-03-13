package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "fmt"
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


func rpsResult(enemy_choose, player_choose string) string {
  if enemy_choose == player_choose {
    return "D"
  } else if enemy_choose == "A"{
    if player_choose == "B" {
      return "W"
    } else {
      return "L"
    }
  } else if enemy_choose == "B" {
    if player_choose == "C" {
      return "W"
    } else {
      return "L"
    }
  } else if enemy_choose == "C" {
    if player_choose == "A" {
      return "W"
    } else {
      return "L"
    }
  }

  return "D"
}


func secondResult(enemy_choose, player_choose string) string{
  if player_choose == "X" {
    // condition must lose
    if enemy_choose == "A" {
      return "C"
    } else if enemy_choose == "B" {
      return "A"
    } else if enemy_choose == "C" {
      return "B" 
    }
  } else if player_choose == "Y" {
    // condition draw
    return enemy_choose
  } else {
    // condition win
    if enemy_choose == "A" {
      return "B"
    } else if enemy_choose == "B" {
      return "C"
    } else if enemy_choose == "C" {
      return "A"
    }
  }

  return "A"
}


func solution() int {
  // init the paper, rock, and scissors point
  var point_shape = map[string]int{
    "A": 1, // Rock
    "B": 2, // Paper
    "C": 3, // Scissors
  }

  var point_result = map[string]int{
    "L": 0, // Lose
    "D": 3, // Draw
    "W": 6, // Win
  }

  var sugested_shape = map[string]string{
    "X": "A",
    "Y": "B",
    "Z": "C",
  }

  var result_point int

  for _, lines := range readFiles() {
    res := strings.Split(lines, " ")

    // res[0] = enemy_choose, res[1] = player_choose
    result_point += point_result[rpsResult(res[0], sugested_shape[res[1]])] + point_shape[sugested_shape[res[1]]]
  }

  return result_point
}


func solution2() int {
  // init the paper, rock, and scissors point
  var point_shape = map[string]int{
    "A": 1, // Rock
    "B": 2, // Paper
    "C": 3, // Scissors
  }

  var point_result = map[string]int{
    "L": 0, // Lose
    "D": 3, // Draw
    "W": 6, // Win
  }

  var sugested_strategy = map[string]string{
    "X": "L",
    "Y": "D",
    "Z": "W",
  }

  var result_point int

  for _, lines := range readFiles() {
    res := strings.Split(lines, " ")

    // res[0] = enemy_choose, res[1] = player_choose
    result_point += point_result[sugested_strategy[res[1]]] + point_shape[secondResult(res[0], res[1])]
  }

  return result_point
}

func main() {
  fmt.Println(solution2())
}
