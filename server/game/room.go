package game

import (
	"errors"
)

type room struct {
	Gui     string `json:"gui"`
	currentPlayers uint8
	maxPlayer uint8
}

const (
	noSpaceLeftInRoom = "this game has reached its max capacity of players"
)

func NewRoom(gui string, maxPlayers int) *room {
	return &room{gui, 0, maxPlayers}
}
