package main

import (
	"fmt"

	_ "sar-server/db"
)

func main() {
	fmt.Println("BC Search and Rescue Field Server")
	go broadcastAnnouncement()
	go runServer()
}
