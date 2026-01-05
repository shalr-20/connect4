//Adding WebSocket Support

package main

import (
	"fmt"
	"net/http"

	"connect4-backend/game"

	"github.com/gorilla/websocket"
)

type ClientMessage struct {
	Type     string `json:"type"`
	Username string `json:"username,omitempty"`
	Column   int    `json:"column,omitempty"`
}

type ServerMessage struct {
	Type   string     `json:"type"`
	Board  game.Board `json:"board,omitempty"`
	Winner int        `json:"winner,omitempty"`
	Turn   int        `json:"turn,omitempty"`
}

var manager = game.NewGameManager()
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	// Read join message
	var joinMsg ClientMessage
	conn.ReadJSON(&joinMsg)

	if joinMsg.Type != "join" {
		return
	}

	// Try reconnect
	activeGame := manager.Rejoin(joinMsg.Username, conn)
	if activeGame != nil {
		conn.WriteJSON(ServerMessage{
			Type:  "reconnected",
			Board: activeGame.State.Board,
			Turn:  activeGame.State.Turn,
		})
		return
	}

	// Otherwise join as new player
	player := game.PlayerConn{
		Username: joinMsg.Username,
		Conn:     conn,
	}

	activeGame = manager.Join(player)

	if activeGame == nil {
		conn.WriteJSON(ServerMessage{
			Type: "waiting",
		})
		select {} // keep connection alive
	}

	conn.WriteJSON(ServerMessage{
		Type:  "game_start",
		Board: activeGame.State.Board,
		Turn:  activeGame.State.Turn,
	})
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("Server running at ws://localhost:8080/ws")
	http.ListenAndServe(":8080", nil)
}
