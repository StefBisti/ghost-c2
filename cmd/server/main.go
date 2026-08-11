package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/StefBisti/ghost-c2/internal/protocol"
)

const listenAddr = "localhost:3000"

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error listening: ", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on ", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			continue
		}
		fmt.Printf("New Connection: %s\n", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	for {
		var beacon protocol.Beacon
		if err := decoder.Decode(&beacon); err != nil {
			if err.Error() == "EOF" {
				fmt.Printf("Agent %s disconnected\n", conn.RemoteAddr())
			} else {
				fmt.Println("Read error:", err)
			}
			return
		}
		fmt.Printf("[%s] Beacon from %s, host: %s\n", time.Now().Format(time.RFC3339), beacon.ID, beacon.Hostname)
	}
}
