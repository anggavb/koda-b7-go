package readfile

import (
	"fmt"
	"io"
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
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	return string(content), nil

	// using os.File
	// content := make([]byte, 1024)
	// n, err := file.Read(content)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// convert byte slice to string cause content is ascii text of bytes
	// return string(content[:n]), nil

	// using bufio if the file is large, we can read it line by line
	// scanner := bufio.NewScanner(file)
	// var result string
	// for scanner.Scan() {
	// 	result += scanner.Text() + "\n"
	// }
	// if err := scanner.Err(); err != nil {
	// 	panic(err.Error())
	// }
	// return result, nil
}