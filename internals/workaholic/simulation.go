package workaholic

import (
	"fmt"
	"time"
)

func Routine() {
	WakeUp()
	MakeTheBed()
	GetOutBed()
	BoilWater()
	Exercise()
	Rest()
	TakeBath()
	WearingClothes()
}

func WakeUp() {
	fmt.Println("Bunyi Alarm... 🔔🚨⏰")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Bangunnn!!! 🐓")
}

func MakeTheBed() {
	fmt.Println("Merapikan tempat tidur...")
	time.Sleep(2 * time.Second)
	fmt.Println("Tempat tidur selesai dirapihkan.")
}

func GetOutBed() {
	fmt.Println("Beranjak dari kamar tidur ke dapur...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Tiba di dapur.")
}

func BoilWater() {
	fmt.Println("Memasak air...")
	time.Sleep(3 * time.Second)
	fmt.Println("Air matang...")
}

func PrepLunch() {
	fmt.Println("Menyiapkan makanan untuk bekal...")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Makanan selesai dibuat.")
}

func MakeCoffee() {
	fmt.Println("Membuat kopi...")
	time.Sleep(1 * time.Second)
	fmt.Println("Selesai membuat kopi.")
}

func Exercise() {
	fmt.Println("Olahraga...")
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai berolahraga.")
}

func Rest() {
	fmt.Println("Beristirahat...")
	time.Sleep(700 * time.Millisecond)
	fmt.Println("Selesai istirahat.")
}

func TakeBath() {
	fmt.Println("Mandi...")
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai Mandi.")
}

func WearingClothes() {
	fmt.Println("Memakai baju...")
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai memakai baju.")
}
