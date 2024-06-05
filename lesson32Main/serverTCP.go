package main

import (
	"bufio"
	"fmt"
	"lessons32Main/file"
	"lessons32Main/model"
	"log"
	"net"
	"strings"
	"sync"
)

var (
	users           []model.User
	usersLock       sync.Mutex
	connections     = make(map[string]net.Conn)
	connectionsLock sync.Mutex
)

func main() {
	// Start a TCP listener on port 8087
	listener, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Println("Server is listening on port 8087")

	file.Read()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Client connected from %s", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	user := model.User{
		IpAddress: conn.RemoteAddr().String(),
	}

	usersLock.Lock()
	users = append(users, user)
	usersLock.Unlock()

	connectionsLock.Lock()
	connections[user.IpAddress] = conn
	connectionsLock.Unlock()

	file.SaveUsers()

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		message = strings.TrimSpace(message)
		log.Printf("Received message: %s", message)
		sendMessageToAll(conn, message)

	}

	log.Printf("client is not connected %s", conn.RemoteAddr().String())

	usersLock.Lock()
	for i := 0; i < len(users); i++ {
		if users[i].IpAddress == conn.RemoteAddr().String() {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	usersLock.Unlock()

	connectionsLock.Lock()
	delete(connections, user.IpAddress)
	connectionsLock.Unlock()

	file.SaveUsers()
}

func sendMessageToAll(senderConn net.Conn, message string) {
	connectionsLock.Lock()
	defer connectionsLock.Unlock()

	for ip, conn := range connections {
		if conn != senderConn {
			_, err := conn.Write([]byte(fmt.Sprintf("Message from %s: %s\n", senderConn.RemoteAddr().String(), message)))
			if err != nil {
				log.Printf("Sending message is error %s: %v", ip, err)
			}
		}
	}
}
