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
func NewSnek(name string, x, y uint) *Snek {
	s := &Snek{
		Name: name,
	}
	s.BodyParts.PushFront(
		&SnekBodyPart{
			X: x,
			Y: y,
		},
	)

	return s
}

func (s *Snek) Move(incrementX, incrementY int) {
	// TODO move like a snake
	// TODO make sure we don't overflow X or Y
	oldX, oldY := s.BodyParts.At(0).X, s.BodyParts.At(0).Y

	if incrementX > 0 {
		s.BodyParts.At(0).X++
	} else if incrementX < 0 {
		s.BodyParts.At(0).X--
	}

	if incrementY > 0 {
		s.BodyParts.At(0).Y++
	} else if incrementY < 0 {
		s.BodyParts.At(0).Y--
	}

	log.Trace().
		Str("name", s.Name).
		Uint("oldX", oldX).
		Uint("oldY", oldY).
		Uint("newX", s.BodyParts.At(0).X).
		Uint("newY", s.BodyParts.At(0).Y).
		Int("incrementX", incrementX).
		Int("incrementY", incrementY).
		Msg("Moved Snek")
}
