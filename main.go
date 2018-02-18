package main

import (
	"context"
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

	ctx := context.Background()

	go broadcastAnnouncement()
	go runServer()

	for {
		select {
		case _, ok := <-ctx.Done():
			if !ok {
				log.Println("[INFO] Exiting through context")
				return
			}
		}
	}
}
