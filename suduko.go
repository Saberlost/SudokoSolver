package main

func validateLineBad(line []int) bool {
	if len(line) != 9 {
		return false
	}
	for i := 1; i < 10; i++ {
		foundOnce := false
		for _, j := range line {
			if i == j {
				if !foundOnce {
					foundOnce = true
				} else {
					return false
				}
			}
		}
		if !foundOnce {
			return false
		}
	}

	return true
}

func validateLine(line []int) bool {
	if len(line) != 9 {
		return false
	}
	var used [10]bool
	nrfound := 0
	for _, j := range line {
		if used[j] {
			return false
		}
		used[j] = true
		nrfound++
	}
	if nrfound == 9 {
		return true
	}
	return false
}

func validateRow(board [][]int, col int) bool {
	col = col
	if len(board) != 9 {
		return false
	}

	var used [10]bool //number already used
	nrfound := 0
	for _, j := range board {
		if used[j[col]] {
			return false
		}
		used[j[col]] = true
		nrfound++
	}
	if nrfound == 9 {
		return true
	}
	return false
}

func getBoxCords(square int) (int, int) {
	lineindex := (square % 3) * 3
	rowindex := square - square%3

	return lineindex, rowindex
}

//Square is 1,2,3
//          4,5,6
//          7,8,9
func validateBox(board [][]int, square int) bool {
	var used [10]bool
	nrfound := 0
	lineindex, rowindex := getBoxCords(square)

	var lineids = []int{lineindex, lineindex + 1, lineindex + 2}
	for _, i := range lineids {
		for j := rowindex; j < rowindex+3; j++ {

			if used[board[j][i]] {
				return false
			}
			used[board[j][i]] = true
			nrfound++
		}
	}
	if nrfound == 9 {
		return true
	}
	return false

}
func isValidPlaceMentLine(line []int, col int, value int) bool {

	if line[col] == value {
		return true
	}
	if line[col] > 0 {
		return false
	}
	for _, v := range line {
		if v == value {
			return false
		}
	}
	return true
}
func isValidPlacementRow(board [][]int, col int, row int, value int) bool {
	if board[row][col] == value {
		return true
	}
	if board[row][col] > 0 {
		//if board[row][col] == value {
		return false
	}
	for _, v := range board {
		if v[col] == value {
			return false
		}
	}
	return true
}

func getSquareFromCords(col int, row int) int {

	//0,1,2  == 0 , 3,4,5 == 1, 6, 7,8 ==2
	//colstart := col / 3

	// 0,1,2 == 0
	// 3,4,5 == 3
	// 6,7,8 == 6
	//rowstart := row - row%3

	square := (row - row%3) + col/3

	return square
}

//IsValidPlaceMentBox test
func IsValidPlaceMentBox(board [][]int, col int, row int, value int) bool {
	square := getSquareFromCords(col, row)

	if board[row][col] == value {
		return true
	}
	if board[row][col] > 0 {

	}

	linestart, rowstart := getBoxCords(square)

	for i := linestart; i < linestart+3; i++ {
		for j := rowstart; j < rowstart+3; j++ {
			if board[j][i] == value {
				return false
			}
		}
	}

	return true
}
