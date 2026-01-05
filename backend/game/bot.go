package game

func copyBoard(board Board) Board {
	var newBoard Board
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			newBoard[r][c] = board[r][c]
		}
	}
	return newBoard
}
func findWinningMove(board Board, player int) int {
	for col := 0; col < Cols; col++ {
		temp := copyBoard(board)
		if DropDisc(&temp, col, player) {
			if CheckWin(temp, player) {
				return col
			}
		}
	}
	return -1
}
func BotMove(board Board) int {
	// 1️⃣ Can bot win?
	if col := findWinningMove(board, Bot); col != -1 {
		return col
	}

	// 2️⃣ Can player win next? Block!
	if col := findWinningMove(board, Player); col != -1 {
		return col
	}

	// 3️⃣ Prefer center columns
	preferred := []int{3, 2, 4, 1, 5, 0, 6}
	for _, col := range preferred {
		temp := copyBoard(board)
		if DropDisc(&temp, col, Bot) {
			return col
		}
	}

	return -1
}
