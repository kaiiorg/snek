package game

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (g *Game) Init() tea.Cmd {
	// Should already be init'd via New
	return nil
}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return g, tea.Quit
		}
	}

	return g, nil
}

func (g *Game) View() string {
	topBottom := strings.Repeat("X", int(g.World.X())) + "\n"
	sides := "X" + strings.Repeat(" ", int(g.World.X()-2)) + "X\n"

	v := topBottom
	for i := uint(0); i < g.World.Y()-2; i++ {
		v += sides
	}
	v += topBottom

	return v
}
