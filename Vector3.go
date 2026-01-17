package main

import "math"

type Vector3 struct {
	x float64
	y float64
	z float64
}

func vector3(x float64, y float64, z float64) Vector3 {
	return Vector3{x, y, z}
}

func (vec *Vector3) AddScalar(scalar float64) Vector3 {
	return vector3(vec.x+scalar, vec.y+scalar, vec.z+scalar)
}

func (vec *Vector3) MultiplyScalar(scalar float64) Vector3 {
	return vector3(vec.x*scalar, vec.y*scalar, vec.z*scalar)
}

func Vector3Zero() Vector3 {
	return vector3(0.0, 0.0, 0.0)
}

func (vec *Vector3) Normalize() Vector3 {
	length := math.Sqrt((vec.x*vec.x + vec.y*vec.y + vec.z*vec.z))
	x := vec.x / length
	y := vec.y / length
	z := vec.z / length
	return Vector3{x, y, z}
}
