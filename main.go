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

	insertDataToSlice()
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
	var triangle string
	
	for i := range input {
		triangle = "*"
		for range i {
			triangle += "*"
		}

		fmt.Println(triangle)
	}
}

func insertDataToSlice() {
	slice := []int{50,75,66,20,32,90}
	newNum := 88

	newSlice := make([]int, 0, 7)
	newSlice = append(newSlice, slice[0:3]...)
	newSlice = append(newSlice, newNum)
	newSlice = append(newSlice, slice[3:]...)

	for _, val := range newSlice {
		fmt.Println(val)
	}
}