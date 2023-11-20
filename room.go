package main

import (
	"github.com/google/uuid"
)

type Spaceship struct {
	XPos          float32
	YPos          float32
	Rotation      float32
	XVelocity     float32
	YVelocity     float32
	Acceleration  float32
	PolarVelocity float32
}

type Missile struct {
	XPos     float32
	YPos     float32
	Rotation float32
	Velocity float32
}

type Player struct {
	Name      string
	Score     int32     `json:"score"`
	Spaceship Spaceship `json:"spaceship"`
}

type Room struct {
	RoomId     uuid.UUID
	Player1    Player
	Player2    Player
	Spaceship1 Spaceship
	Spaceship2 Spaceship
	Missiles1  []Missile
	Missiles2  []Missile
}

func NewRoom(roomId uuid.UUID) *Room {
	var player1 *Player
	player1 = new(Player)
	player1.Name = "Player1"
	player1.Score = 0

	var player2 *Player
	player2 = new(Player)
	player2.Name = "Player1"
	player2.Score = 0

	var room *Room
	room = new(Room)
	room.RoomId = roomId
	room.Player1 = *player1
	room.Player2 = *player2

	var spaceship1 *Spaceship
	spaceship1 = new(Spaceship)
	room.Spaceship1 = *spaceship1

	var spaceship2 *Spaceship
	spaceship2 = new(Spaceship)
	room.Spaceship1 = *spaceship2

	// room.Missiles1 =
	return room
}

type UpdatePlayerRequest struct {
	Player    Player
	Spaceship Spaceship
	Missiles  []Missile
}

type UpdatePlayerResponse struct {
	Room Room
}

// type Manager struct {
// 	roomChannels map[string]broadcast.Broadcaster
// 	open         chan *Listener
// 	close        chan *Listener
// 	delete       chan string
// 	messages     chan *Message
// }

// func NewRoomManager() *Manager {
// 	manager := &Manager{
// 		roomChannels: make(map[string]broadcast.Broadcaster),
// 		open:         make(chan *Listener, 100),
// 		close:        make(chan *Listener, 100),
// 		delete:       make(chan string, 100),
// 		messages:     make(chan *Message, 100),
// 	}

// 	go manager.run()
// 	return manager
// }
