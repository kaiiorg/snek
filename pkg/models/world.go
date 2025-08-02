package models

import "time"

// World keeps track of everything in the world of Snek
type World struct {
	// tick is the current step of the world. Used to help keep track of things as they change. Incremented each game tick.
	tick uint64
	// tickDelay is roughly how long is between each game tick.
	tickDelay time.Duration

	// playerSnek points to the local player's snek
	// If this is null, the game is running in headless mode
	playerSnek *Snek
	// allSneks points to all sneks loaded into the world
	allSneks []*Snek

	// boundaryX sets the rightmost limit of the world. A snek contacting 0 or this limit dies a painful death
	boundaryX uint
	// boundaryY sets the uppermost limit of the world. A snek contacting 0 or this limit dies a painful death
	boundaryY uint
}

// NewWorld initializes a World with a given tickDelay and size
func NewWorld(tickDelay time.Duration, x, y uint) *World {
	w := &World{
		tickDelay: tickDelay,
		allSneks:  []*Snek{},
		boundaryX: x,
		boundaryY: y,
	}

	return w
}

func (w *World) X() uint {
	return w.boundaryX
}

func (w *World) Y() uint {
	return w.boundaryY
}
