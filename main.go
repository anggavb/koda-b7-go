package main

import (
	"fmt"
	"math"
)

type Education struct {
	name string
	major string
}

type Biodata struct {
	name string
	photo string
	email string
	age uint
	telp string
	isMarriage bool
	education []Education
}

func main() {
	// Minitask 1
	radius := 12
    area, circumference := calculateCircle(radius)
	
	fmt.Printf("Jari-jari: %d\nLuas: %.2f\nKeliling: %.2f\n", radius, area, circumference)
	
	// Minitask 2
	input := 5
	triangleFromTriangle(input)

	// Minitask 3
	insertDataToSlice()

	// Minitask 4
	angga := Biodata{
		name: "Angga Vb",
		photo: "image.png",
		email: "anggavb8@gmail.com",
		age: 28,
		telp: "0851 5677 0131",
		isMarriage: true,
		education: []Education{
			{name: "SMK AL-BAHRI BEKASI", major: "Rekayasa Perangkat Lunak"},
			{name: "Politeknik LP3I Cileungsi", major: "Information Management"},
		},
	}

	fmt.Println(angga)
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