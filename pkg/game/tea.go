package game

import (
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

	rendered := g.Renderer.Render()
	g.PreviousRender = &rendered

	return rendered
}

func (g *Game) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "w", "up":
		// Move up
		g.IncrementX.Store(0)
		g.IncrementY.Store(-1)
	case "a", "left":
		// Move left
		g.IncrementX.Store(-1)
		g.IncrementY.Store(0)
	case "s", "down":
		// Move down
		g.IncrementX.Store(0)
		g.IncrementY.Store(1)
	case "d", "right":
		// Move right
		g.IncrementX.Store(1)
		g.IncrementY.Store(0)
	case "ctrl+c":
		return g, tea.Quit
	default:
		log.Trace().Str("keyMesg", msg.String()).Msg("Unhandled key input")
	}
	return g, nil
}
