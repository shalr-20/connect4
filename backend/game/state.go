package game

type GameState struct {
	Board  Board
	Turn   int
	Status string // "playing", "finished"
	Winner string
}
func NewGame() *GameState {
	return &GameState{
		Board:  NewBoard(),
		Turn:   Player,
		Status: "playing",
	}
}
