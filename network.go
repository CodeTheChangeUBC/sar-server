package main

import (
	"net"
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

	addr := net.UDPAddr{
		IP:   BroadcastIPv4,
		Port: Port,
	}

	for {
		_, err := socket.WriteToUDP([]byte(IdentifierMessage), &addr)

		if err != nil {
			go broadcastAnnouncement()
			panic(err)
		}

		time.Sleep(5 * time.Second)
	}
}
