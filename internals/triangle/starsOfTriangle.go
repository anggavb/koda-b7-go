package triangle

import "fmt"

func StarsFromTriangle(input int) {
	var triangle string
	
	for i := range input {
		triangle = "*"
		for range i {
			triangle += "*"
		}

		fmt.Println(triangle)
	}
}