package data

import "fmt"

func InsertDataToSlice() {
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