package readfile

import (
	"os"
)

// Minitask 6: Read File Content
func ReadFile(filename string) (string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer file.Close()

	// Read the file content
	content := make([]byte, 1024)
	n, err := file.Read(content)
	if err != nil {
		return "", err
	}
	return string(content[:n]), nil
}