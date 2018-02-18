package main

import (
	"fmt"
	"log"

	sdb "sar-server/db"
)

func main() {
	fmt.Println("BC Search and Rescue Field Server")

	path := "./sar.db"
	db, err := sdb.InitializeDB(path)
	if err != nil {
		log.Fatalf("Failed to open database %s. Error returned: %e", path, err)
	}

	sdb.Database = db

	go broadcastAnnouncement()
	go runServer()
}
