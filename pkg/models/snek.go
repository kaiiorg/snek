package models

import "github.com/gammazero/deque"

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
