package renderer

import (
	"strings"
	"time"

	"github.com/kaiiorg/snek/pkg/models"

	"github.com/rs/zerolog/log"
)

var (
	space    = ' '
	lf       = '\n'
	wall     = 'X'
	snekDead = '∩'
	snekHead = 'S'
	snekBody = 's'
)

type Renderer struct {
	world *models.World

	cachedTopBottom []rune
	cachedSides     []rune

	currentFrame []rune

	stringBuilder strings.Builder

	metricsStart      time.Time
	metricsFrameInit  time.Duration
	metricsSnekRender time.Duration
	metricsRuneRender time.Duration
}

func NewRenderer(world *models.World) *Renderer {
	r := &Renderer{
		world:           world,
		cachedTopBottom: make([]rune, world.X()+1),
		cachedSides:     make([]rune, world.X()+1),
	}

	r.initTopBottom()
	r.initSides()

	return r
}

func (r *Renderer) Render() string {
	r.initFrame()
	r.renderSneks()
	result := r.renderFrameToString()

	log.Trace().
		Str("frameInit", r.metricsFrameInit.String()).
		Str("snekRender", r.metricsSnekRender.String()).
		Str("runeRender", r.metricsRuneRender.String()).
		Str("total", (r.metricsFrameInit + r.metricsSnekRender + r.metricsRuneRender).String()).
		Msg("Render time")

	return result
}
