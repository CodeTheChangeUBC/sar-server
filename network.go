package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

const (
	// Port to connect to.
	Port int = 6543

	// IdentifierMessage is the message we broadcast out.
	IdentifierMessage string = "sarcnc"
)

// BroadcastIPv4 address.
var BroadcastIPv4 = net.IPv4(255, 255, 255, 255)

func broadcastAnnouncement() {
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   BroadcastIPv4,
		Port: Port,
	})

	if err != nil {
		go broadcastAnnouncement()
		panic(err)
	}

	for {
		_, err = socket.Write([]byte(IdentifierMessage))

		if err != nil {
			go broadcastAnnouncement()
			panic(err)
		}

		time.Sleep(5 * time.Second)
	}
}

func server() {
	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.FormatInt(int64(Port), 10))
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		sock, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %e", err)
		}

		go handleClient(sock)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
}
