package madding

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var boards []BoardMessage

type BoardMessage struct {
	Name, Message string
}

func Sender(msgChan chan string, name *string) {
	// for range msgChan {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Masukkan nama anda: ")
		scanner.Scan()
		*name = strings.TrimSpace(scanner.Text())
		fmt.Print("Masukkan pesan anda: ")
		scanner.Scan()
		message := strings.TrimSpace(scanner.Text())
	
		msgChan	<- message
	// }
}

func Board(msg, name string) {
	boards = append(boards, BoardMessage{Name: name, Message: msg})

	fmt.Println("-------- Notice Board --------")
	for _, board := range boards {
		fmt.Printf("Name: %s   |   Message: %s\n", board.Name, board.Message)
	}
}