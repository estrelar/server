package player

import (
	m "../model"
)

type Player struct {
	Name     string
	Position m.Position
}

func (p *Player) ToProxy() PlayerProxy {
	return PlayerProxy{
		Name:     p.Name,
		Position: p.Position,
	}
}
