import { useEffect, useState } from "react";

export default function Leaderboard() {
  const [data, setData] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/leaderboard")
      .then((res) => res.json())
      .then(setData);
  }, []);

  return (
    <div>
      <h3>🏅 Leaderboard</h3>
      <ul>
        {data.map((u) => (
          <li key={u.username}>
            {u.username} - {u.wins} wins
          </li>
        ))}
      </ul>
    </div>
  );
}
