package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello World test")
	fmt.Println("what should i write here")
	board := getBoard("testfile.txt")

	start := time.Now()

	solved, counter := solveBoard(board, 0, 0)
	elapsed := time.Since(start)
	if solved {
		fmt.Println("board solved in just ", counter, " steps")
		for _, row := range board {
			fmt.Println("", row)
		}
	}

	log.Printf("suduko solver took %s", elapsed)

}

func solveBoard(board [][]int, row int, col int) (bool, int) {
	counter := 0
	if col > 8 {
		row++
		col = 0
	}
	if row > 8 {
		//fmt.Println("This should be 9,9 and solved board: ", board)
		return true, counter
	}

	for value := 1; value < 10; value++ {

		if isValidPlaceMentLine(board[row], col, value) {
			if isValidPlacementRow(board, col, row, value) {
				if IsValidPlaceMentBox(board, col, row, value) {
					counter++
					//if row == 0 && col > 0 {
					//fmt.Println(value, "isValid for ", row, ",", col, ":", board)
					//}
					oldvalue := board[row][col]
					board[row][col] = value
					ok, dcounter := solveBoard(board, row, col+1)
					counter += dcounter
					if ok {

						return true, counter
					}
					board[row][col] = oldvalue

				}
			}
		}
	}
	return false, counter
}

func getBoard(filename string) [][]int {

	file, err := os.Open(filename)
	if err != nil {
		panic(1)
	}

	scn := bufio.NewScanner(file)

	var col, row int
	col = 0
	row = 0
	//var board [][]int
	board := make([][]int, 9)

	for scn.Scan() {
		col = 0
		data := scn.Text()
		board[row] = make([]int, 9)

		for _, i := range strings.Split(data, " ") {

			if col > 8 {
				panic(1)
			}
			board[row][col], _ = strconv.Atoi(i)

			col++
		}
		row++

	}
	for _, row := range board {
		fmt.Println("", row)
	}
	return board
}
