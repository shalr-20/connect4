package game

const (
	Rows = 6
	Cols = 7

	Empty  = 0
	Player = 1
	Bot    = 2
)

type Board [Rows][Cols]int
func NewBoard() Board {
	var board Board
	return board
}
func DropDisc(board *Board, col int, player int) bool {
	if col < 0 || col >= Cols {
		return false
	}

	for row := Rows - 1; row >= 0; row-- {
		if board[row][col] == Empty {
			board[row][col] = player
			return true
		}
	}
	return false
}
func CheckWin(board Board, player int) bool {
	// Horizontal
	for r := 0; r < Rows; r++ {
		for c := 0; c <= Cols-4; c++ {
			if board[r][c] == player &&
				board[r][c+1] == player &&
				board[r][c+2] == player &&
				board[r][c+3] == player {
				return true
			}
		}
	}

	// Vertical
	for c := 0; c < Cols; c++ {
		for r := 0; r <= Rows-4; r++ {
			if board[r][c] == player &&
				board[r+1][c] == player &&
				board[r+2][c] == player &&
				board[r+3][c] == player {
				return true
			}
		}
	}

	// Diagonal ↘
	for r := 0; r <= Rows-4; r++ {
		for c := 0; c <= Cols-4; c++ {
			if board[r][c] == player &&
				board[r+1][c+1] == player &&
				board[r+2][c+2] == player &&
				board[r+3][c+3] == player {
				return true
			}
		}
	}

	// Diagonal ↙
	for r := 3; r < Rows; r++ {
		for c := 0; c <= Cols-4; c++ {
			if board[r][c] == player &&
				board[r-1][c+1] == player &&
				board[r-2][c+2] == player &&
				board[r-3][c+3] == player {
				return true
			}
		}
	}

	return false
}
func IsDraw(board Board) bool {
	for c := 0; c < Cols; c++ {
		if board[0][c] == Empty {
			return false
		}
	}
	return true
}
