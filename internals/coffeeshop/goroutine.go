package coffeeshop

import (
	"fmt"
	"time"
)

func Barista(id int, menu string) {
	fmt.Printf("Pesanan %s sedang dibuat oleh Barista %d\n", menu, id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Pesanan %s selesai\n", menu)
}