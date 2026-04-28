package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/anggavb/koda-b7-go/internals/circle"
	"github.com/anggavb/koda-b7-go/internals/data"
	"github.com/anggavb/koda-b7-go/internals/models"
	"github.com/anggavb/koda-b7-go/internals/readfile"
	"github.com/anggavb/koda-b7-go/internals/triangle"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n------- Main Menu -------")
		fmt.Println("1. Calculate the Area & Circumference of a Circle")
		fmt.Println("2. Generate Right-Angled Triangle")
		fmt.Println("3. Inserting New Number to Slice")
		fmt.Println("4. Show My Biodata")
		fmt.Println("5. Read File Content")
		fmt.Println("6. Input Person")
		fmt.Println("0. Exit")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			fmt.Print("Insert radius of circle: ")
			if !scanner.Scan() {
				break
			}
			radiusInput := strings.TrimSpace(scanner.Text())
			var radius int
			fmt.Sscanf(radiusInput, "%d", &radius)

			area, circumference := circle.CalculateCircle(radius)
			fmt.Printf("Jari-jari: %d\nLuas: %.2f\nKeliling: %.2f\n", radius, area, circumference)

		case "2":
			fmt.Print("Insert the number of rows for the triangle: ")
			if !scanner.Scan() {
				break
			}
			rowsInput := strings.TrimSpace(scanner.Text())
			var rows int
			fmt.Sscanf(rowsInput, "%d", &rows)

			triangle.StarsFromTriangle(rows)

		case "3":
			fmt.Print("Insert a new number to be added to the slice: ")
			if !scanner.Scan() {
				break
			}
			newNumInput := strings.TrimSpace(scanner.Text())
			var newNum int
			fmt.Sscanf(newNumInput, "%d", &newNum)
			data.InsertDataToSlice(newNum)

		case "4":
			myBiodata()

		case "5":
			fmt.Print("Insert the filename to read: ") // dir/file.txt
			if !scanner.Scan() {
				break
			}
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()
			filename := strings.TrimSpace(scanner.Text())
			content, err := readfile.ReadFile(filename)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("File Content:")
			fmt.Println(content)

		case "6":
			fmt.Print("Insert person name: ")
			if !scanner.Scan() {
				break
			}
			name := strings.TrimSpace(scanner.Text())
			
			fmt.Print("Insert person address: ")
			if !scanner.Scan() {
				break
			}
			address := strings.TrimSpace(scanner.Text())
			
			fmt.Print("Insert person phone: ")
			if !scanner.Scan() {
				break
			}
			phone := strings.TrimSpace(scanner.Text())

			person := models.NewPerson(name, address, phone)
			fmt.Println("Data Person Berhasil Dibuat!")

			fmt.Println(person.Print())
			fmt.Println(person.Greet())
			person.SetNamePerson("upin")
			fmt.Println(person.Greet())

		default:
			fmt.Println("Thanks 👋")
			return
		}
	}
}

func myBiodata() {
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