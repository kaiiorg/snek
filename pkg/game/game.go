package game

import (
	"context"
	"flag"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rs/zerolog/log"
	"sync/atomic"
	"time"

	"github.com/kaiiorg/snek/pkg/models"
)

var (
	tickDelay = flag.Duration("tick-delay", time.Second, "Roughly how long between each game tick")
	worldX    = flag.Uint("x", 120, "How wide the world should be")
	worldY    = flag.Uint("y", 29, "How tall the world should be")
	headless  = flag.Bool("headless", false, "If no player snek should be spawned")
)

// Game performs actions on a World and the sneks in it
type Game struct {
	SkipRender     atomic.Bool
	World          *models.World
	PreviousRender *string
}

func New() *Game {
	g := &Game{
		World: models.NewWorld(*tickDelay, *worldX, *worldY),
	}

	centerX, centerY := g.World.Center()
	if !*headless {
		g.World.SpawnSnek("player", centerX, centerY, true)
	}

	return g
}

func (g *Game) Run() {
	ctx, ctxCancel := context.WithCancel(context.Background())

	p := tea.NewProgram(g)
	go func() {
		defer ctxCancel()
		if _, err := p.Run(); err != nil {
			log.Error().Err(err).Msg("Tea Exploded")
		}
	}()

	t := time.NewTicker(*tickDelay)
	for {
		select {
		case <-ctx.Done():
			log.Warn().Msg("Snek exiting")
			return
		case <-t.C:
			g.tick(p)
		}
	}
}

func (g *Game) tick(teaProgram *tea.Program) {
	// TODO do game logic

	// Tell the render method it is ok to actually do its thing now
	g.SkipRender.Store(false)
	teaProgram.Send(true)
}
