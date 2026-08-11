package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/StefBisti/ghost-c2/internal/protocol"
)

const (
	serverAddress = "localhost:3000"
	interval      = 10 * time.Second
)

func main() {
	conn := connect()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	beaconSignal(&conn)
	for range ticker.C {
		beaconSignal(&conn)
	}
}

func beaconSignal(conn *net.Conn) {
	_, err := sendMessage(*conn)
	if err != nil {
		fmt.Println("Lost connection, reconnecting...")
		(*conn).Close()
		(*conn) = connect()
	}
}

func connect() net.Conn {
	for {
		conn, err := net.Dial("tcp", serverAddress)
		if err != nil {
			fmt.Println("Connection failed, retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}
		fmt.Println("Connected to server")
		return conn
	}
}

func sendMessage(conn net.Conn) (int, error) {
	beacon := protocol.Beacon{ID: "agent-001", Hostname: "kali"}
	data, err := json.Marshal(beacon)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal beacon: %w", err)
	}
	return conn.Write(data)
}
