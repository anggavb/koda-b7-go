package readfile

import (
	"fmt"
	"os"
)

// Minitask 6: Read File Content
func ReadFile(filename string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered while reading file:", r)
		}
	}()
    file, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer file.Close()

	// Read the file content
	content := make([]byte, 1024)
	n, err := file.Read(content)
	if err != nil {
		panic(err.Error())
	}
	// convert byte slice to string cause content is ascii text of bytes
	return string(content[:n]), nil
}