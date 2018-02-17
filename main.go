package main

import (
	"fmt"
)

func main() {
	fmt.Println("BC Search and Rescue Field Server")
	go broadcastAnnouncement()
	go runServer()
}
