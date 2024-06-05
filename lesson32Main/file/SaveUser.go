package file

import (
	"encoding/json"
	"lessons32Main/model"
	"log"
	"net"
	"os"
	"sync"
)

var (
	users           []model.User
	usersLock       sync.Mutex
	connections     = make(map[string]net.Conn)
	connectionsLock sync.Mutex
)

func SaveUsers() {
	usersLock.Lock()
	defer usersLock.Unlock()

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Failed to open users file: %v", err)
	}
	defer file.Close()

	data := json.NewEncoder(file)
	err = data.Encode(users)
	if err != nil {
		log.Fatalf("Failed to encode users: %v", err)
	}
}

func Read() {
	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open users file: %v", err)
	}
	defer file.Close()

	data1 := json.NewDecoder(file)
	err = data1.Decode(&users)
	if err != nil {
		log.Printf("No existing users or failed to decode: %v", err)
	}
}
