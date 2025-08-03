package renderer

import (
	"time"

	"github.com/kaiiorg/snek/pkg/models"
)

func (r *Renderer) initFrame() {
	// TODO write tests
	r.metricsStart = time.Now()
	r.currentFrame = r.currentFrame[:0] // Keep backing array allocated
	r.currentFrame = append(r.currentFrame, r.cachedTopBottom...)
	for i := uint(0); i < r.world.Y()-2; i++ {
		r.currentFrame = append(r.currentFrame, r.cachedSides...)
	}
	r.currentFrame = append(r.currentFrame, r.cachedTopBottom...)
	r.metricsFrameInit = time.Since(r.metricsStart)
}

func (r *Renderer) renderFrameToString() string {
	// TODO write tests
	r.metricsStart = time.Now()
	r.stringBuilder.Reset()
	for i := range r.currentFrame {
		r.stringBuilder.WriteRune(r.currentFrame[i])
	}
	result := r.stringBuilder.String()
	r.metricsRuneRender = time.Since(r.metricsStart)
	return result
}

func (r *Renderer) renderSneks() {
	// TODO write tests
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
				r.currentFrame[r.xyCoordsTo1DOffset(part.X, part.Y)] = symbol
			}
		}
		r.metricsSnekRender = time.Since(r.metricsStart)
	})
}
