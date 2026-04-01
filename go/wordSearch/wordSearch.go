package main

func wordSearch(board [][]byte, word string) bool {
	var firstLetterCoord [][]int

	// Find all coordinates that has the initial letter
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				firstLetterCoord = append(firstLetterCoord, []int{i, j})
			}
		}
	}

	if len(firstLetterCoord) < 1 {
		return false
	}

	// For each coordinate of first letter, check if it is possible to complete word
	for _, v := range firstLetterCoord {
		found := checkNextLetter(board, v, 1, word)

		if found {
			return true
		}
	}

	return false
}

func checkNextLetter(board [][]byte, coord []int, currentLetterIndex int, word string) bool {
	y, x := coord[0], coord[1]
	left, right, above, below := []int{y, x - 1}, []int{y, x + 1}, []int{y - 1, x}, []int{y + 1, x}
	checks := [][]int{left, right, above, below}

	for _, k := range checks {

		// Coordinate out of range
		if k[0] < 0 || k[1] < 0 || k[0] >= len(board) || k[1] >= len(board[0]) {
			continue
		}

		if board[k[0]][k[1]] == word[currentLetterIndex] {
			if currentLetterIndex == len(word)-1 {
				return true
			}

			currentLetterIndex++
			return checkNextLetter(board, []int{k[0], k[1]}, currentLetterIndex, word)
		}

	}

	return false
}
