package main

import "math"

type Vector2Int struct {
	X int32
	Y int32
}

func (v Vector2Int) Add(other Vector2Int) Vector2Int {
	return Vector2Int{v.X + other.X, v.Y + other.Y}
}

func (v Vector2Int) Sub(other Vector2Int) Vector2Int {
	return Vector2Int{v.X - other.X, v.Y - other.Y}
}

func (v Vector2Int) Mul(other Vector2Int) Vector2Int {
	return Vector2Int{v.X * other.X, v.Y * other.Y}
}

func (v Vector2Int) Div(other Vector2Int) Vector2Int {
	return Vector2Int{v.X / other.X, v.Y / other.Y}
}

func (v Vector2Int) Scale(scalar int32) Vector2Int {
	return Vector2Int{v.X * scalar, v.Y * scalar}
}

func (v Vector2Int) Length() float32 {
	return float32(v.X*v.X + v.Y*v.Y)
}

func (v Vector2Int) Normalize() Vector2Int {
	return v.Scale(int32(1 / v.Length()))
}

func (v Vector2Int) Dot(other Vector2Int) int32 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector2Int) Equals(other Vector2Int) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v Vector2Int) AngleDeg() float64 {
	return math.Atan2(float64(v.Y), float64(v.X)) * 180 / math.Pi
}

func (v Vector2Int) Cross(other Vector2Int) float32 {
	return float32(v.X)*float32(other.Y) - float32(v.Y)*float32(other.X)
}
