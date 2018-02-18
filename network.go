package main

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	// Port to connect to.
	Port int = 6543

	// MagicByte is included in every frame for the TCP protocol
	MagicByte uint8 = 0x5B

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
			continue
		}

		go handleClient(sock)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4096)
	frame, err := readFrame(conn, buf)
}

// A Frame for our TCP protocol.
type Frame struct {
	Magic   byte
	Length  uint32
	Payload []byte
}

// When an error is returned from this function the stream is in an undefined
// state. It's recommended that an error message be sent and the connection
// closed.
func readFrame(conn net.Conn, buf []byte) (Frame, error) {
	read, err := conn.Read(buf[0:5])

	if read != 5 {
		panic("Something's rotten in the state of Denmark")
	} else if err != nil {
		return Frame{}, err
	}

	var frame Frame

	frame.Magic = buf[0]
	frame.Length = binary.LittleEndian.Uint32(buf[1:5])

	// SECURITY PROBLEM. Easy DoS.
	frame.Payload = make([]byte, frame.Length)

	totalRead := uint32(0)

	for totalRead != frame.Length {
		read, err := conn.Read(frame.Payload[totalRead:])
		if err != nil {
			return Frame{}, err
		}

		totalRead += uint32(read)
	}

	return frame, nil
}
