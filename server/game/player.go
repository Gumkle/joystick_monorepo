package game

type Player struct {
	Nickname string
	localId unit8
}

func (p *Player) SendMessage(message int) {
	p.room.Chan <- message + uint8
}