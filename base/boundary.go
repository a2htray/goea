package base

import "github.com/a2htray/goea/math"

// Boundary
type Boundary math.Vector

func (b Boundary) Vector() math.Vector {
	return math.Vector(b)
}

// Limit contains the limitation of the upper and lower boundary
type Limit struct {
	Upper Boundary
	Lower Boundary
}