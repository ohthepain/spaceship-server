package main

import (
	"fmt"

	"github.com/google/uuid"
)

type State struct {
	Rooms map[uuid.UUID]Room
}

func NewState() *State {
	fmt.Printf("NewState\n")
	var s State
	s.Rooms = make(map[uuid.UUID]Room)
	var roomid = uuid.New()
	s.Rooms[roomid] = *NewRoom(roomid)
	return &s
}

var state State = *NewState()
