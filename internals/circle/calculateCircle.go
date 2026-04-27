package circle

import "math"

func CalculateCircle(r int) (area float32, circumference float32) {
	var pi float32 = math.Phi
	area = calculateCircleArea(r, pi)
	circumference = calculateCircleCircumference(r, pi)

	return area, circumference
}

func calculateCircleArea(r int, pi float32) (area float32) {
	area = pi * (float32(r) * float32(r))
	return area
}

func calculateCircleCircumference(r int, pi float32) (circumference float32) {
	circumference = float32(2) * pi * float32(r)
	return circumference
}