package db

import (
	"context"
	"log"

	"github.com/PUDDLEEE/puddleee_back/ent"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *ent.Client {
	client, err := ent.Open("sqlite3", "file:test.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
