package utils

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v Vec2) Normalize() Vec2 {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vec2{v.X / magnitude, v.Y / magnitude}
}
