package game

import (
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"time"

	"github.com/kaiiorg/snek/pkg/models"
)

var (
	tickDelay = flag.Duration("tick-delay", time.Second, "Roughly how long between each game tick")
	worldX    = flag.Uint("x", 120, "How wide the world should be")
	worldY    = flag.Uint("y", 29, "How tall the world should be")
)

// Game performs actions on a World and the sneks in it
type Game struct {
	World *models.World
}

func New() *Game {
	g := &Game{
		World: models.NewWorld(*tickDelay, *worldX, *worldY),
	}

	return g
}

func (g *Game) Run() {
	p := tea.NewProgram(g)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Tea exploded: %s\n", err)
		os.Exit(1)
	}
}
