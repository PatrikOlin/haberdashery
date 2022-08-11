package main

import (
	"log"

	"github.com/PatrikOlin/haberdashery/pkg/db"
)

func init() {
	_, err := db.Open()
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
}

func main() {
	clearOrphans()
}

func clearOrphans() {
	stmt := "DELETE FROM garments WHERE is_orphan = true"

	_, err := db.DBClient.Exec(stmt)
	if err != nil {
		log.Fatalln(err)
	}
}
