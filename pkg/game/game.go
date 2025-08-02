package game

import (
	"flag"
	"fmt"
	"time"

	"github.com/kaiiorg/snek/pkg/models"
)

var (
	tickDelay = flag.Duration("tick-delay", time.Second, "Roughly how long between each game tick")
	worldX    = flag.Uint("x", 30, "How wide the world should be")
	worldY    = flag.Uint("y", 20, "How tall the world should be")
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
	fmt.Printf("Game run\n")
}
