package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to the server at localhost:8080
	conn, err := net.Dial("tcp", "localhost:8087")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to the server")

	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)
	// Create a reader for server response
	serverReader := bufio.NewReader(conn)

	for {
		// Read user input
		fmt.Print("Enter message: ")
		userMessage, _ := reader.ReadString('\n')
		userMessage = strings.TrimSpace(userMessage)

		// Send user input to the server
		_, err = fmt.Fprintf(conn, "%s\n", userMessage)
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}

		// Read server response
		serverMessage, err := serverReader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read server response: %v", err)
		}
		fmt.Printf("Server response: %s", serverMessage)

	}
}
