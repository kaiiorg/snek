package models

import (
	"github.com/gammazero/deque"
	"github.com/kaiiorg/snek/pkg/tools"
	"github.com/rs/zerolog/log"
)

// Snek is a player character. May be a local or remote character.
type Snek struct {
	// Name is this snek's name. Will be displayed where appropriate.
	Name string

	// Dead is if this snek ran into something and died. How sad.
	Dead bool

	// BodyParts describes all parts of a snek
	BodyParts deque.Deque[*SnekBodyPart]

	// clipper handles the logic for clipping sneks and objects to within the world boundary
	clipper *tools.Clipper
}

type SnekBodyPart struct {
	X uint
	Y uint
}

// NewSnek initializes a Snek and sets its location to the requested position
func NewSnek(name string, x, y, length uint, clipper *tools.Clipper) *Snek {
	s := &Snek{
		Name:    name,
		clipper: clipper,
	}

	// Must have at least a head
	s.BodyParts.PushFront(&SnekBodyPart{X: x, Y: y})
	// Add however many body segments we need
	for i := uint(0); i < length; i++ {
		y--
		s.BodyParts.PushFront(&SnekBodyPart{X: x, Y: y})
	}

	return s
}

func (s *Snek) Move(incrementX, incrementY int) {
	// TODO make sure we don't overflow X or Y
	oldX, oldY := s.BodyParts.At(0).X, s.BodyParts.At(0).Y
	newX, newY := oldX, oldY

	moved := false
	if incrementX > 0 {
		moved = true
		newX++
	} else if incrementX < 0 {
		moved = true
		newX--
	}

	if incrementY > 0 {
		moved = true
		newY++
	} else if incrementY < 0 {
		moved = true
		newY--
	}

	if !moved {
		return
	}

	// If we are more than just a head, make sure we can't backtrack over our body
	if s.BodyParts.Len() > 1 {
		bodyX, bodyY := s.BodyParts.At(1).X, s.BodyParts.At(1).Y
		if bodyX == newX && bodyY == newY {
			return
		}
	}

	// Determine if we've collided with the world edge
	// TODO Not thread safe to update the bool this way
	newX, newY, clipped := s.clipper.World(newX, newY)

	// If we were clipped to the world boundary, don't update the position and mark this snek as dead
	if clipped {
		s.Dead = clipped
	} else {
		oldTail := s.BodyParts.PopBack()
		oldTail.X = newX
		oldTail.Y = newY
		s.BodyParts.PushFront(oldTail)
	}

	log.Trace().
		Str("name", s.Name).
		Uint("oldX", oldX).
		Uint("oldY", oldY).
		Uint("newX", newY).
		Uint("newY", newY).
		Bool("clipped", clipped).
		Bool("dead", s.Dead).
		Int("incrementX", incrementX).
		Int("incrementY", incrementY).
		Msg("Moved Snek")
}
