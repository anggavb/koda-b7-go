package main

import (
	"fmt"

	"github.com/anggavb/koda-b7-go/internals/circle"
	"github.com/anggavb/koda-b7-go/internals/data"
	"github.com/anggavb/koda-b7-go/internals/models"
	"github.com/anggavb/koda-b7-go/internals/triangle"
)

func main() {
	// Minitask 1
	radius := 12
    area, circumference := circle.CalculateCircle(radius)
	
	fmt.Printf("Jari-jari: %d\nLuas: %.2f\nKeliling: %.2f\n", radius, area, circumference)
	
	// Minitask 2
	input := 5
	triangle.StarsFromTriangle(input)

	// Minitask 3
	data.InsertDataToSlice()

	// Minitask 4
	angga := models.Biodata{
		Name: "Angga Vb",
		Photo: "image.png",
		Email: "anggavb8@gmail.com",
		Age: 28,
		Telp: "0851 5677 0131",
		IsMarriage: true,
		Education: []models.Education{
			{Name: "SMK AL-BAHRI BEKASI", Major: "Rekayasa Perangkat Lunak"},
			{Name: "Politeknik LP3I Cileungsi", Major: "Information Management"},
		},
	}

	fmt.Println(angga)
}