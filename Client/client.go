package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Welcome,", name)
	fmt.Println("Type messages below. Type 'exit' to leave.")

	var lastSeen int
	var mu sync.Mutex

	go func() {
		for {
			var messages []string
			err := client.Call("ChatServer.GetMessages", "", &messages)
			if err != nil {
				fmt.Println("[Server disconnected]")
				time.Sleep(time.Second)
				continue
			}

			mu.Lock()
			if len(messages) > lastSeen {
				newMsgs := messages[lastSeen:]
				for _, msg := range newMsgs {
					if strings.HasPrefix(msg, name+":") {
						text := strings.TrimPrefix(msg, name+":")
						text = strings.TrimSpace(text)
						fmt.Printf("you: %s\n", text)
					} else {
						fmt.Println(msg)
					}
				}
				lastSeen = len(messages)
			}
			mu.Unlock()

			time.Sleep(800 * time.Millisecond)
		}
	}()

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		if text == "exit" {
			fmt.Println("Leaving chat...")
			break
		}

		message := fmt.Sprintf("%s: %s", name, text)
		var reply string
		err := client.Call("ChatServer.SendMessage", message, &reply)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		mu.Lock()
		lastSeen++
		mu.Unlock()
	}
}
