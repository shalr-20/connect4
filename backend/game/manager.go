package game

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type PlayerConn struct {
	Username     string
	Conn         *websocket.Conn
	Disconnected bool
}

type ServerMessage struct {
	Type   string       `json:"type"`
	Board  Board   `json:"board,omitempty"`
	Winner string          `json:"winner,omitempty"`
	Turn   int          `json:"turn,omitempty"`
}

type ActiveGame struct {
	ID          string
	Player1     PlayerConn
	Player2     PlayerConn
	State       *GameState
	LastMoveAt  time.Time
	Mutex       sync.Mutex
}


type GameManager struct {
	WaitingPlayer *PlayerConn
	Games         map[string]*ActiveGame
	Mutex         sync.Mutex
}
func NewGameManager() *GameManager {
	return &GameManager{
		Games: make(map[string]*ActiveGame),
	}
}
func (gm *GameManager) Join(player PlayerConn) *ActiveGame {
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	// If someone is waiting → create PvP game
	if gm.WaitingPlayer != nil {
		game := &ActiveGame{
			ID: time.Now().Format("20060102150405"),
			Player1: *gm.WaitingPlayer,
			Player2: player,
			State:   NewGame(),
		}
		gm.Games[game.ID] = game
		gm.WaitingPlayer = nil
		return game
	}

	// Otherwise wait
	gm.WaitingPlayer = &player

	// Bot fallback after 10s
	go func() {
		time.Sleep(10 * time.Second)
		gm.Mutex.Lock()
		defer gm.Mutex.Unlock()

		if gm.WaitingPlayer != nil && gm.WaitingPlayer.Username == player.Username {
			game := &ActiveGame{
				ID: time.Now().Format("20060102150405"),
				Player1: *gm.WaitingPlayer,
				Player2: PlayerConn{Username: "BOT"},
				State:   NewGame(),
			}
			gm.Games[game.ID] = game
			gm.WaitingPlayer = nil
		}
	}()

	return nil
}
func (g *ActiveGame) HandleDisconnect(player *PlayerConn) {
	player.Disconnected = true
	g.LastMoveAt = time.Now()

	go func() {
		time.Sleep(30 * time.Second)
		g.Mutex.Lock()
		defer g.Mutex.Unlock()

		if player.Disconnected && g.State.Status == "playing" {
			// Forfeit the game
			var winner string
			if player.Username == g.Player1.Username {
				winner = g.Player2.Username
			} else {
				winner = g.Player1.Username
			}
			g.State.Status = "finished"
			g.State.Winner = winner

			// Notify remaining player
			var remainingPlayer *PlayerConn
			if player.Username == g.Player1.Username {
				remainingPlayer = &g.Player2
			} else {
				remainingPlayer = &g.Player1
			}

			if remainingPlayer.Conn != nil {
				remainingPlayer.Conn.WriteJSON(ServerMessage{
					Type:   "game_over",
					Board:  g.State.Board,
					Winner: winner,
				})
			}
		}
	}()
}
func (gm *GameManager) Rejoin(username string, conn *websocket.Conn) *ActiveGame {
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	for _, g := range gm.Games {
		if g.State.Status == "playing" {
			if g.Player1.Username == username {
				g.Player1.Conn = conn
				g.Player1.Disconnected = false
				return g
			}
			if g.Player2.Username == username {
				g.Player2.Conn = conn
				g.Player2.Disconnected = false
				return g
			}
		}
	}
	return nil
}
