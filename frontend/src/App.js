import { useState } from "react";
import Game from "./Game";
import Leaderboard from "./Leaderboard";

function App() {
  const [username, setUsername] = useState("");
  const [joined, setJoined] = useState(false);

  return (
    <div style={{ padding: 20 }}>
      <h2>🎯 Connect 4</h2>

      {!joined ? (
        <>
          <input
            placeholder="Enter username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <button onClick={() => setJoined(true)}>Join Game</button>
        </>
      ) : (
        <>
          <Game username={username} />
          <Leaderboard />
        </>
      )}
    </div>
  );
}

export default App;
