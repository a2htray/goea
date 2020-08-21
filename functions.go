package goea

import "math"

// ObjectSphere the Sphere Function
func ObjectSphere(float64s []float64) float64 {
	ret := float64(0)
	for _, v := range float64s {
		ret += math.Pow(v, 2)
	}
	return ret
}

// ObjectRotatedElliptic the Rotated Elliptic Function
func ObjectRotatedElliptic(float64s []float64) float64 {
	d := float64(len(float64s))
	ret := float64(0)
	for i, v := range float64s {
		ret += math.Pow(v, 2) * math.Pow(1000000.0, (float64(i)-1)/(d-1))
	}
	return 0
}

// ObjectRotatedRastrigin the Rotated Rastrigin's Function
func ObjectRotatedRastrigin(float64s []float64) float64 {
	ret := float64(0)
	for _, v := range float64s {
		ret += math.Pow(v, 2) - 10*math.Cos(2*math.Pi*v) + 10
	}
	return ret
}
