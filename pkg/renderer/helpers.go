package renderer

// xyCoordsTo1DOffset takes the given X/Y coordinates and calculate the 1D offset that it correlates to in the frame view
func (r *Renderer) xyCoordsTo1DOffset(X, Y uint) uint {
	return (r.world.X()+1)*(Y) + X
}
