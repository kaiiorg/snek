package game

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rs/zerolog/log"
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
	topBottom := strings.Repeat("X", int(g.World.X())) + "\n"
	sides := "X" + strings.Repeat(" ", int(g.World.X()-2)) + "X\n"

	render := topBottom
	for i := uint(0); i < g.World.Y()-2; i++ {
		render += sides
	}
	render += topBottom
	log.Info().Msg("World rendered")

	g.PreviousRender = &render

	return render
}

func (g *Game) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return g, tea.Quit
	}
	return g, nil
}
