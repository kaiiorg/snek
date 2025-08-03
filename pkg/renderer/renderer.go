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

	// TODO shove this logic into a method and write a test for it
	// Build our cached top/bottom
	// All Xs, then an LF
	// Example: XXXXXXX\n
	for i := range r.cachedTopBottom {
		r.cachedTopBottom[i] = wall
	}
	r.cachedTopBottom[len(r.cachedTopBottom)-1] = lf

	// TODO shove this logic into a method and write a test for it
	// Build our cached sides
	// All spaces, with Xs on both side, then an LF
	// Example: X     X\n
	for i := range r.cachedSides {
		r.cachedSides[i] = space
	}
	r.cachedSides[0] = wall
	r.cachedSides[len(r.cachedSides)-2] = wall
	r.cachedSides[len(r.cachedSides)-1] = lf

	return r
}

func (r *Renderer) Render() string {
	// TODO shove this logic into a method and write a test for it
	r.metricsStart = time.Now()
	var frame []rune
	frame = append(frame, r.cachedTopBottom...)
	for i := uint(0); i < r.world.Y()-2; i++ {
		frame = append(frame, r.cachedSides...)
	}
	frame = append(frame, r.cachedTopBottom...)
	r.metricsFrameInit = time.Since(r.metricsStart)

	r.world.RenderSneks(func(sneks []*models.Snek) {
		r.metricsStart = time.Now()
		for _, snek := range sneks {
			first := true
			symbol := snekHead
			for part := range snek.BodyParts.Iter() {
				// Determine which symbol to use; head, dead, or body
				switch {
				case !first:
					symbol = snekBody
				case first && !snek.Dead:
					first = false
				case first && snek.Dead:
					first = false
					symbol = snekDead
				}
				frame[r.findOffset(part.X, part.Y)] = symbol
			}
		}
		r.metricsSnekRender = time.Since(r.metricsStart)
	})

	// TODO shove this logic into a method and write a test for it
	r.metricsStart = time.Now()
	r.stringBuilder.Reset()
	for i := range frame {
		r.stringBuilder.WriteRune(frame[i])
	}
	result := r.stringBuilder.String()
	r.metricsRuneRender = time.Since(r.metricsStart)

	log.Trace().
		Str("frameInit", r.metricsFrameInit.String()).
		Str("snekRender", r.metricsSnekRender.String()).
		Str("runeRender", r.metricsRuneRender.String()).
		Str("total", (r.metricsFrameInit + r.metricsSnekRender + r.metricsRuneRender).String()).
		Msg("Render time")

	return result
}

func (r *Renderer) findOffset(X, Y uint) uint {
	return (r.world.X()+1)*(Y) + X
}
