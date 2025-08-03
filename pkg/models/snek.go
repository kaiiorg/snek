package models

import (
	"github.com/gammazero/deque"
	"github.com/rs/zerolog/log"
)

// Snek is a player character. May be a local or remote character.
type Snek struct {
	// Name is this snek's name. Will be displayed where appropriate.
	Name string

	// BodyParts describes all parts of a snek
	BodyParts deque.Deque[*SnekBodyPart]
}

type SnekBodyPart struct {
	X uint
	Y uint
}

// NewSnek initializes a Snek and sets its location to the requested position
func NewSnek(name string, x, y, length uint) *Snek {
	s := &Snek{
		Name: name,
	}

	s.BodyParts.PushFront(&SnekBodyPart{X: x, Y: y})
	for i := uint(0); i < length; i++ {
		x--
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

	oldTail := s.BodyParts.PopBack()
	oldTail.X = newX
	oldTail.Y = newY
	s.BodyParts.PushFront(oldTail)

	log.Trace().
		Str("name", s.Name).
		Uint("oldX", oldX).
		Uint("oldY", oldY).
		Uint("newX", newY).
		Uint("newY", newY).
		Int("incrementX", incrementX).
		Int("incrementY", incrementY).
		Msg("Moved Snek")
}
