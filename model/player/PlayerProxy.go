package player

import (
	m "../model"
)

type PlayerProxy struct {
	Name     string     `json:"name"`
	Position m.Position `json:"position"`
}
