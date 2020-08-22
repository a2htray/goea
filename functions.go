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
		ret += math.Pow(v, 2) * math.Pow(1000000.0, (float64(i))/(d-1))
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

// ObjectHyperEllipsoid Hyper-Ellipsoid
func ObjectHyperEllipsoid(float64s []float64) float64 {
	ret := float64(0)
	for i, v := range float64s {
		ret += math.Pow(float64(i+1), 2) * math.Pow(v, 2)
	}
	return ret
}

// ObjectKatsuura Katsuura's Function
func ObjectKatsuura(beta int) func([]float64) float64 {
	return func(float64s []float64) float64 {
		ret := float64(1)

		for i, v := range float64s {
			j := float64(i + 1)
			part := float64(0)

			for k := 0; k <= beta; k++ {
				a := math.Pow(2, float64(k)) * v
				part += math.Abs(a-float64(FindNearestInt(a))) * math.Pow(2, -float64(k))
			}

			ret *= 1 + j*part
		}

		return ret
	}
}

// ObjectGriewangk Griewangk’s function
func ObjectGriewangk(float64s []float64) float64 {
	a := float64(0)
	for _, v := range float64s {
		a += math.Pow(v, 2)
	}
	a = a / 4000

	b := float64(1)
	for i, v := range float64s {
		b *= math.Cos(v / (float64(i) + 1))
	}

	return a - b + 1
}
