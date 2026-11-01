# 🎯 4 in a Row – Real-Time Multiplayer Game

A real-time, backend-driven implementation of **4 in a Row (Connect Four)** built using **GoLang**, **WebSockets**, **React**, **PostgreSQL**, and **Kafka (bonus analytics)**.

This project was developed as part of a **Backend Engineering Intern Assignment**, focusing on real-time systems, concurrency, game logic, persistence, and event-driven architecture.

---

## 🧠 Features

### 🎮 Gameplay
- Classic **7×6 Connect Four** board
- Turn-based real-time gameplay
- Win detection:
  - Horizontal
  - Vertical
  - Diagonal
- Draw detection when the board is full

### 👥 Multiplayer & Bot
- Player vs Player matchmaking
- **10-second wait time** for opponent
- Automatic **competitive bot fallback**
- Bot strategy:
  - Tries to win when possible
  - Blocks opponent’s immediate win
  - Prefers center columns

### 🌐 Real-Time System
- WebSocket-based communication
- Backend-authoritative game state
- Instant board updates after every move

### 🔁 Reconnection Handling
- Players can reconnect within **30 seconds**
- Game state preserved during disconnect
- Automatic forfeit if player fails to reconnect

### 🧾 Persistence
- Completed games stored in **PostgreSQL**
- Player win count tracked
- Persistent leaderboard

### 🏅 Leaderboard
- Displays number of wins per player
- Sorted by highest wins
- Available via frontend and REST API

### 💥 Bonus: Kafka Analytics
- Game events published to Kafka
- Separate analytics consumer
- Tracks:
  - Average game duration
  - Most frequent winners
  - Games per hour/day

---

## 🏗️ Architecture Overview
Frontend (React)

▼

Backend (Go)

▼

Kafka Consumer (Analytics Service)


---

## ⚙️ Setup Instructions

### 1️⃣ Prerequisites
- Go 1.22+
- Node.js 18+
- PostgreSQL
- Kafka (optional, for analytics)

---

### 2️⃣ Backend Setup

```
cd backend
go mod tidy
go run main.go
```

Frontend setup:
```
cd frontend
npm install
npm start
