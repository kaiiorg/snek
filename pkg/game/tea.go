package game

import (
	"bytes"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiiorg/snek/pkg/models"
	"github.com/rs/zerolog/log"
)

var (
	empty    = []byte{}
	space    = []byte(" ")
	lf       = []byte("\n")
	wall     = []byte("X")
	snekHead = []byte("S")[0]
	snekBody = []byte("s")[0]
)

func (g *Game) Init() tea.Cmd {
	// Should already be init'd via New
	return nil
}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return g.handleKeyMsg(msg)
	}

	return g, nil
}

func (g *Game) View() string {
	// Don't render unless the game logic told us to do so.
	// This saves us from rendering on every single tea message
	if g.SkipRender.Load() {
		return *g.PreviousRender
	}
	g.SkipRender.Store(true)

	// TODO move render logic to its own struct
	topBottom := bytes.Join(
		[][]byte{
			bytes.Repeat(wall, int(g.World.X())),
			lf,
		},
		empty,
	)
	sides := bytes.Join(
		[][]byte{
			wall,
			bytes.Repeat(space, int(g.World.X()-2)),
			wall,
			lf,
		},
		empty,
	)

	// Build a full frame out of the top/bottom and sides
	render := bytes.Join(
		[][]byte{
			topBottom,
			bytes.Repeat(sides, int(g.World.Y()-2)),
			topBottom,
		},
		empty,
	)

	log.Info().Msg("World frame rendered")

	g.World.RenderSneks(func(sneks []*models.Snek) {
		for i, snek := range sneks {
			first := true
			for part := range snek.BodyParts.Iter() {
				if first {
					render[g.renderFindOffset(part.X, part.Y)] = snekHead
				} else {
					render[g.renderFindOffset(part.X, part.Y)] = snekBody
				}

			}
			log.Info().Int("snek", i).Msg("Rendered snek")
		}
	})

	renderStr := string(render)
	g.PreviousRender = &renderStr

	return renderStr
}

func (g *Game) renderFindOffset(X, Y uint) uint {
	return (g.World.X()+1)*(Y-1) + X
}

func (g *Game) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return g, tea.Quit
	}
	return g, nil
}
