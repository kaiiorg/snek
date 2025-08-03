package models

import (
	"sync"
	"time"

	"github.com/kaiiorg/snek/pkg/tools"

	"github.com/rs/zerolog/log"
)

// World keeps track of everything in the world of Snek
type World struct {
	// tick is the current step of the world. Used to help keep track of things as they change. Incremented each game tick.
	tick uint64
	// tickDelay is roughly how long is between each game tick.
	tickDelay time.Duration

	// playerSnek points to the local player's snek
	// If this is null, the game is running in headless mode
	playerSnek   *Snek
	playerSnekMu sync.RWMutex
	// allSneks points to all sneks loaded into the world
	allSneks   []*Snek
	allSneksMu sync.RWMutex

	// boundaryX sets the rightmost limit of the world. A snek contacting 0 or this limit dies a painful death
	boundaryX uint
	// boundaryY sets the uppermost limit of the world. A snek contacting 0 or this limit dies a painful death
	boundaryY uint

	// clipper handles the logic for clipping sneks and objects to within the world boundary
	clipper *tools.Clipper
}

// NewWorld initializes a World with a given tickDelay and size
func NewWorld(tickDelay time.Duration, x, y uint) *World {
	if x < 20 {
		x = 20
	}
	if y < 10 {
		y = 10
	}

	w := &World{
		tickDelay: tickDelay,
		allSneks:  []*Snek{},
		boundaryX: x,
		boundaryY: y,
		clipper: &tools.Clipper{
			WorldX: x,
			WorldY: y,
		},
	}

	return w
}

func (w *World) X() uint {
	return w.boundaryX
}

func (w *World) Y() uint {
	return w.boundaryY
}

func (w *World) Center() (uint, uint) {
	return w.boundaryX / 2, w.boundaryY / 2
}

func (w *World) SpawnSnek(name string, x, y uint, player bool) {
	// Set X/Y to center if either are out of range
	if x >= w.boundaryX || y >= w.boundaryY {
		x, y = w.Center()
	}

	s := NewSnek(name, x, y, 5, w.clipper)
	if player {
		w.playerSnekMu.Lock()
		w.playerSnek = s
		w.playerSnekMu.Unlock()
	}
	w.allSneksMu.Lock()
	w.allSneks = append(w.allSneks, s)
	w.allSneksMu.Unlock()

	log.Info().
		Str("name", name).
		Uint("x", x).
		Uint("y", y).
		Bool("player", player).
		Msg("Snek spawned")
}

func (w *World) RenderSneks(renderer func(sneks []*Snek)) {
	// TODO this may result in a deadlock if we're not careful. May need to tweak this in the future.
	w.playerSnekMu.RLock()
	defer w.playerSnekMu.RUnlock()
	w.allSneksMu.RLock()
	defer w.allSneksMu.RUnlock()
	renderer(w.allSneks)
}

func (w *World) UpdatePlayerSnek(incrementX, incrementY int) {
	if w.playerSnek == nil {
		return
	}
	w.playerSnekMu.Lock()
	w.playerSnek.Move(incrementX, incrementY)
	w.playerSnekMu.Unlock()

	log.Debug().
		Int("x", incrementX).
		Int("y", incrementY).
		Msg("Moved player snek")
}
