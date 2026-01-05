import { useEffect, useRef, useState } from "react";

const ROWS = 6;
const COLS = 7;

function emptyBoard() {
  return Array.from({ length: ROWS }, () =>
    Array(COLS).fill(0)
  );
}

export default function Game({ username }) {
  const ws = useRef(null);
  const [board, setBoard] = useState(emptyBoard());
  const [turn, setTurn] = useState(1);
  const [status, setStatus] = useState("");

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/ws");

    ws.current.onopen = () => {
      ws.current.send(
        JSON.stringify({ type: "join", username })
      );
    };

    ws.current.onmessage = (e) => {
      const msg = JSON.parse(e.data);

      if (msg.board) setBoard(msg.board);
      if (msg.turn) setTurn(msg.turn);

      if (msg.type === "waiting") setStatus("Waiting for opponent...");
      if (msg.type === "game_start") setStatus("Game started!");
      if (msg.type === "game_over")
        setStatus(msg.winner === 1 ? "You won!" : "You lost!");
      if (msg.type === "draw") setStatus("Draw!");
    };

    return () => ws.current.close();
  }, [username]);

  function playMove(col) {
    ws.current.send(
      JSON.stringify({ type: "move", column: col })
    );
  }

  return (
    <div>
      <h3>{status}</h3>

      <div style={{ display: "grid", gridTemplateColumns: "repeat(7, 40px)" }}>
        {board.map((row, r) =>
          row.map((cell, c) => (
            <div
              key={`${r}-${c}`}
              onClick={() => playMove(c)}
              style={{
                width: 40,
                height: 40,
                border: "1px solid black",
                textAlign: "center",
                fontSize: 24,
                cursor: "pointer",
              }}
            >
              {cell === 1 ? "🔴" : cell === 2 ? "🟡" : ""}
            </div>
          ))
        )}
      </div>
    </div>
  );
}
