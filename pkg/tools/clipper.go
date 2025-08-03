package tools

// Clipper abstracts away logic used to increment and clip X/Y coordinates to various planes
type Clipper struct {
	WorldX uint
	WorldY uint
}

// World clips the given X/Y coordinates to within the WorldX and WorldY coordinates
// and note if it was clipped (aka, collided)
func (c *Clipper) World(newX, newY uint) (uint, uint, bool) {
	newX, clippedX := c.clip(newX, c.WorldX)
	newY, clippedY := c.clip(newY, c.WorldY)
	return newX, newY, clippedX || clippedY
}

// clip clips the given 1D coordinate to within the world coordinate and notes if it was clipped (aka, collided)
// The new value must be [1, world-2] (inclusive)
func (c *Clipper) clip(new, world uint) (uint, bool) {
	// If went too far left/up, you collided with the world edge
	if new == 0 {
		return 1, true
	}
	// If went too far right/down, you collided with the world edge
	if new >= world-1 {
		return world - 1, true
	}
	// Otherwise, you're free to collide with something else
	return new, false
}
