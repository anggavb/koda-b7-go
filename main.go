package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 12
    area, circumference := calculateCircle(radius)

	fmt.Printf("Jari-jari: %d\nLuas: %.2f\nKeliling: %.2f\n", radius, area, circumference)
	
	input := 5
	triangleFromTriangle(input)
}

func calculateCircle(r int) (area float32, circumference float32) {
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

func triangleFromTriangle(input int) {
	for i := range input {
		var triangle string = "*"
		for range i {
			triangle += "*"
		}

		fmt.Println(triangle)
	}
}