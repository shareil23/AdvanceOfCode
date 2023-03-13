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
	var (
		i_tail, i_head, j_tail, j_head int
		tail_history_move              = make(map[string]int)
	)

	// init first move [0][0] or position s
	tail_history_move["[0][0]"] = 1

	for _, lines := range readFiles() {
		// line_data [0] to tell where the direction
		// line_data [1] to tell step need to take
		line_data := strings.Split(lines, " ")
		total_step, _ := strconv.Atoi(line_data[1])

		switch line_data[0] {
		case "R":
			// for right move
			for i := 0; i < total_step; i++ {
				j_head += 1

				if !isTailMove(i_tail, j_tail, i_head, j_head) {
					continue
				}

				i_tail = i_head
				j_tail = j_head - 1

				// check the key is exists
				if _, is_ok := tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)]; is_ok {
					tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)] += 0
				}

				tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)] = 1
			}
		case "L":
			// for left move
			for i := 0; i < total_step; i++ {
				j_head -= 1

				if !isTailMove(i_tail, j_tail, i_head, j_head) {
					continue
				}

				i_tail = i_head
				j_tail = j_head + 1

				// check the key is exists
				if _, is_ok := tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)]; is_ok {
					tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)] += 0
				}

				tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)] = 1
			}
		case "U":
			// for up move
			for i := 0; i < total_step; i++ {
				i_head += 1

				if !isTailMove(i_tail, j_tail, i_head, j_head) {
					continue
				}

				i_tail = i_head - 1
				j_tail = j_head

				// check the key is exists
				if _, is_ok := tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)]; is_ok {
					tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)] += 0
				}

				tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)] = 1
			}
		case "D":
			// for down move
			for i := 0; i < total_step; i++ {
				i_head -= 1

				if !isTailMove(i_tail, j_tail, i_head, j_head) {
					continue
				}

				i_tail = i_head + 1
				j_tail = j_head

				// check the key is exists
				if _, is_ok := tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)]; is_ok {
					tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)] += 0
				}

				tail_history_move[fmt.Sprintf("[%d][%d]", i_tail, j_tail)] = 1
			}
		}
	}

	return len(tail_history_move)
}

func isTailMove(i_tail, j_tail, i_head, j_head int) bool {
	if (i_tail == i_head && j_tail == j_head) ||
		(i_tail == i_head && j_tail+1 == j_head) ||
		(i_tail == i_head && j_tail-1 == j_head) {
		return false
	} else if (i_tail+1 == i_head && j_tail+1 == j_head) ||
		(i_tail+1 == i_head && j_tail == j_head) ||
		(i_tail+1 == i_head && j_tail-1 == j_head) {
		return false
	} else if (i_tail-1 == i_head && j_tail-1 == j_head) ||
		(i_tail-1 == i_head && j_tail == j_head) ||
		(i_tail-1 == i_head && j_tail+1 == j_head) {
		return false
	}

	return true
}

type coord struct {
	x int
	y int
}

func solution2() int {
	// code here
	var tail_history_move = make(map[coord]bool)
	var knots = make([]coord, 10)

	//init first start value
	tail_history_move[knots[9]] = true

	for _, lines := range readFiles() {
		// line_data [0] to tell where the direction
		// line_data [1] to tell step need to take
		line_data := strings.Split(lines, " ")
		total_step, _ := strconv.Atoi(line_data[1])

		for total_step > 0 {
			switch line_data[0] {
			case "U":
				knots[0].y++
			case "R":
				knots[0].x++
			case "D":
				knots[0].y--
			case "L":
				knots[0].x--
			}

			// move rest of the knot
			for i := range knots[:len(knots)-1] {
				knots[i+1] = adjustTail(knots[i+1], knots[i])
			}

			total_step--
			tail_history_move[knots[9]] = true
		}
	}

	return len(tail_history_move)
}

func adjustTail(tail coord, head coord) coord {
	var newTail coord = tail
	var diff coord = coord{head.x - tail.x, head.y - tail.y} // if head and tail are touching

	switch diff {
	case coord{-2, 1}, coord{-1, 2}, coord{0, 2}, coord{1, 2}, coord{2, 1}, coord{2, 2}, coord{-2, 2}:
		newTail.y++
	}

	switch diff {
	case coord{1, 2}, coord{2, 1}, coord{2, 0}, coord{2, -1}, coord{1, -2}, coord{2, 2}, coord{2, -2}:
		newTail.x++
	}

	switch diff {
	case coord{-2, -2}, coord{2, -1}, coord{1, -2}, coord{0, -2}, coord{-1, -2}, coord{-2, -1}, coord{2, -2}:
		newTail.y--
	}

	switch diff {
	case coord{-2, -2}, coord{-1, -2}, coord{-2, -1}, coord{-2, -0}, coord{-2, 1}, coord{-1, 2}, coord{-2, 2}:
		newTail.x--
	}
	return newTail
}

func main() {
	fmt.Println(solution2())
}
