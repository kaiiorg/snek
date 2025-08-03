package renderer

func (r *Renderer) initTopBottom() {
	// TODO write tests
	// Build our cached top/bottom
	// All Xs, then an LF
	// Example: XXXXXXX\n
	for i := range r.cachedTopBottom {
		r.cachedTopBottom[i] = wall
	}
	r.cachedTopBottom[len(r.cachedTopBottom)-1] = lf
}

func (r *Renderer) initSides() {
	// TODO write tests
	// Build our cached sides
	// All spaces, with Xs on both side, then an LF
	// Example: X     X\n
	for i := range r.cachedSides {
		r.cachedSides[i] = space
	}
	r.cachedSides[0] = wall
	r.cachedSides[len(r.cachedSides)-2] = wall
	r.cachedSides[len(r.cachedSides)-1] = lf
}
