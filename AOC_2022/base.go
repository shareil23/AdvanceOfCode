package main

import (
	"bufio"
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


func solution() {
  // code here
}


func main() {
  // for _, lines := range readFiles() {
      // the code here
  // }
  // fmt.Println(solution())
}
